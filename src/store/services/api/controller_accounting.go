package api

import (
	. "store/models"
	"store/services/gdrv"
	"fmt"
	"time"
	"github.com/asdine/storm/q"
	"github.com/asdine/storm"
	"github.com/cznic/mathutil"
	"github.com/rs/zerolog/log"
)

type ControllerAccounting struct {
	FileID string //Папка в которой храним все накладные
	Controller
}

func (p *ControllerAccounting) IndexGET() {

}

/**
	Поиск заказов где AccountingID == nil
	и заказ неравен OrderStatusDeclined и OrderStatusRefunded
 */
func (p *ControllerAccounting) CheckDeliveryReports() error {
	//каталог
	store := p.GetStore()
	//открыть транзакцию
	tx, err := store.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var ordersBucket = tx.From(NodeNamedOrders)
	var reportsBucket = tx.From(NodeNamedAccounting)

	matcher := q.Eq("AccountingID", 0)

	matcher = q.And(matcher,
		q.Not(q.Eq("Status", OrderStatusDeclined)),
		q.Not(q.Eq("Status", OrderStatusRefunded)),
		q.Not(q.Eq("Status", OrderStatusAwaitingPayment)),
	)

	var orders []Order
	err = ordersBucket.Select(matcher).Find(&orders)
	if err == storm.ErrNotFound {
		return nil
	}

	if err != nil {
		return err
	}

	for _, order := range orders {
		var positions []PositionReport

		for _, position := range order.Positions {
			positions = append(positions, PositionReport{
				ProductSKU: position.ProductSKU,
				Amount: position.Amount,
				Price: position.Total,
			})
		}

		report := AccountingReport{
			Type: ReportTypeDelivery,
			UpdatedAt: time.Now(),
			CreatedAt: order.CreatedAt,
			Positions: positions,
			Source: SourceReportStore,
			SourceID: order.Invoice,
		}

		err := reportsBucket.Save(&report)
		if err != nil {
			return err
		}

		order.AccountingID = report.ID
		err = ordersBucket.Save(&order)
		if err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}

//Проверяем наличие новых отчетов на google drive
func (p *ControllerAccounting) CheckReportsInGoogleDrv() error {
	files, err := gdrv.Walk(p.FileID)

	if err != nil {
		return err
	}

	var catalog = p.GetStore().From(NodeNamedAccounting)

	tx, err := catalog.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, file := range files {
		//по имени получаем дату(18.01.18)
		date, err := time.Parse("02.01.06", file.Name)
		if err != nil {
			fmt.Printf("[Accounting] Error parsing name file: %v", err)
			continue
		}

		//поиск по базе наличие у нас этого файла
		//Source = SourceReportGoogle
		//SourceID = file.Name
		sheets, err := gdrv.GetAllSheets(file.Id)
		if err != nil {
			return err
		}

		for _, sheet := range sheets {
			sheetName := sheet.Properties.Title
			matcher := q.And(
				q.Eq("Source", SourceReportGoogle),
				q.Eq("SourceID", fmt.Sprintf("%s:%s", file.Name, sheetName)))

			report := AccountingReport{
				Type: ReportTypeReceiving,
				Positions: []PositionReport{},
				CreatedAt: date,
				UpdatedAt: time.Now(),
				Source: SourceReportGoogle,
				SourceID: fmt.Sprintf("%s:%s", file.Name, sheetName),
			}
			err = catalog.Select(matcher).First(&report)

			if err != nil && err != storm.ErrNotFound {
				return err
			}

			var rows []SheetAccounting
			err = gdrv.UnmarshalSpreadsheet(&rows, file.Id, sheetName)
			if err != nil {
				return err
			}

			var positions []PositionReport

			for _, sheetData := range rows {
				if len(sheetData.ProductSKU) > 0 {
					positions = append(positions, CreatePositionReport(sheetData))
				}
			}

			report.Positions = positions

			err = tx.Save(&report)
			if err != nil {
				return err
			}
		}
	}

	tx.Commit()

	return nil
}

//Весь список
func (p *ControllerAccounting) GetAllReports(pagination Pagination) (*PageAccountingReport, error) {
	var reports []AccountingReport
	catalog := p.GetStore().From(NodeNamedAccounting)
	err := catalog.AllByIndex("ID", &reports, storm.Limit(pagination.Limit), storm.Skip(pagination.Offset), storm.Reverse())
	if err != nil {
		return nil, err
	}

	total, err := catalog.Count(new(AccountingReport))
	if err != nil {
		return nil, err
	}

	return &PageAccountingReport{
		Content: reports,
		Cursor: Cursor{
			Total:  total,
			Limit:  pagination.Limit,
			Offset: pagination.Offset,
		},
	}, nil
}

//Поиск
func (p *ControllerAccounting) SearchReportsWithFilter(filter FilterAccountingReport, pagination Pagination) (*PageAccountingReport, error) {
	matcher := q.True()

	if len(filter.Type) != 0 {
		matcher = q.And(matcher, q.Eq("Type", filter.Type))
	}

	switch filter.Where {
	case FilterReportWhereRangeDate:
		startDate := filter.StartDate.Truncate(24 * time.Hour)
		endDate := filter.EndDate.AddDate(0, 0, 1).Truncate(24 * time.Hour)
		matcher = q.And(matcher, q.Gte("CreatedAt", startDate), q.Lte("CreatedAt", endDate))
	case FilterReportWhereDate:
		beginningDay := filter.StartDate.Truncate(24 * time.Hour)
		nextDay := filter.StartDate.AddDate(0, 0, 1).Truncate(24 * time.Hour)
		matcher = q.And(matcher, q.Gte("CreatedAt", beginningDay), q.Lte("CreatedAt", nextDay))
	case FilterReportWhereSKU:
		matcher = q.And(matcher, q.Eq("Position.ProductSKU", filter.Query))
	}

	var reports []AccountingReport
	err := p.GetStore().
		From(NodeNamedAccounting).
		Select(matcher).
		Limit(pagination.Limit).
		Skip(pagination.Offset).
		OrderBy("CreatedAt").
		Reverse().
		Find(&reports)

	if err != nil && err != storm.ErrNotFound {
		return nil, err
	}

	total, err := p.GetStore().
		From(NodeNamedAccounting).
		Select(matcher).
		Count(new(AccountingReport))

	if err != nil {
		return nil, err
	}

	return &PageAccountingReport{
		Content: reports,
		Cursor: Cursor{
			Total:  total,
			Limit:  pagination.Limit,
			Offset: pagination.Offset,
		},
	}, nil
}

//Детальная информация
func (p *ControllerAccounting) GetReportByID(id int) (AccountingReport, error) {
	var report AccountingReport
	err := p.GetStore().
		From(NodeNamedAccounting).
		One("ID", id, &report)

	if err != nil {
		return report, err
	}

	return report, nil
}

//Обновить количество остатков
//получаем список всех накладных
//Производим аудит все товаров
//количество приходных - расходных - заказов которые ждут оплаты
func (p *ControllerAccounting) UpdateQuantity() error {
	var store = p.GetStore()

	tx, err := store.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var ordersBucket = tx.From(NodeNamedOrders)
	var catalogBucket = tx.From(NodeNamedCatalog)
	var accountingBucket = tx.From(NodeNamedAccounting)

	var reports []AccountingReport
	err = accountingBucket.AllByIndex("ID", &reports)
	if err != nil && err != storm.ErrNotFound {
		return err
	}

	var holdOrders []Order
	err = ordersBucket.Find("Status", OrderStatusAwaitingPayment, &holdOrders)
	if err != nil && err != storm.ErrNotFound {
		return err
	}

	var products []Product
	err = catalogBucket.AllByIndex("ID", &products)
	if err != nil && err != storm.ErrNotFound {
		return err
	}

	//Удержано
	quantityHold := func(productSKU string) int {
		result := 0

		for _, order := range holdOrders {
			for _, position := range order.Positions {
				if position.ProductSKU == productSKU {
					result = result + position.Amount
				}
			}
		}
		return result
	}

	//Накладные
	quantityReport := func(productSKU string) (received int, delivery int) {
		for _, report := range reports {
			for _, position := range report.Positions {
				if position.ProductSKU == productSKU && report.Type == ReportTypeReceiving {
					received = received + position.Amount
				}
				if position.ProductSKU == productSKU && report.Type == ReportTypeDelivery {
					delivery = delivery + position.Amount
				}
			}
		}
		return
	}

	for _, product := range products {
		hold := quantityHold(product.SKU)
		received, delivery := quantityReport(product.SKU)

		//подсчет количество
		quantity := received - delivery - hold
		if quantity < 0 {
			log.Error().Err(fmt.Errorf("неверный подсчет количества у продукта с артикул %s", product.SKU))
		}
		product.Quantity = mathutil.Max(quantity, 0)

		err := catalogBucket.Save(&product)
		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}

//Очищаем отчеты
func (p *ControllerAccounting) ClearReports() error {
	var store = p.GetStore()

	tx, err := store.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var accountingBucket = tx.From(NodeNamedAccounting)
	var ordersBucket = tx.From(NodeNamedOrders)

	err = accountingBucket.Drop(&AccountingReport{})

	//найти все заказы и установить accounting = nil
	matcher := q.Not(q.Eq("AccountingID", nil))

	var orders []Order
	err = ordersBucket.Select(matcher).Find(&orders)

	if err != nil && err != storm.ErrNotFound {
		return err
	}

	for _, order := range orders {
		order.AccountingID = 0
		err := ordersBucket.Save(&order)
		if err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}