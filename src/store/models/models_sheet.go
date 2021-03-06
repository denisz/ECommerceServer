package models

import (
	"fmt"
	"strings"
	"strconv"
)

type (
	SheetProduct struct {
		Quantity      int      `sheet:"Количество"`
		Price         int      `sheet:"Цена"`
		Weight        int      `sheet:"Вес"`
		Name          string   `sheet:"Наименование"`
		CollectionSKU string   `sheet:"Категория"`
		Producer      string   `sheet:"Производитель"`
		Factor        string   `sheet:"Фасовка"`
		Form          string   `sheet:"Форма"`
		SKU           string   `sheet:"Артикул"`
		Discount      string   `sheet:"Скидка"`
		Pictures      []string `sheet:"Картинки"`
		Dimension     string   `sheet:"ШxДxВ"`
		Disabled      bool     `sheet:"Снят с продажи"`
	}

	SheetProductMedia struct {
		SKU      string   `sheet:"Артикул"`
		Pictures []string `sheet:"Картинки"`
	}

	SheetCollection struct {
		SKU  string `sheet:"Артикул"`
		Name string `sheet:"Наименование"`
	}

	SheetNotation struct {
		SKU         string `sheet:"Артикул"`
		Effects     string `sheet:"Эффекты"`
		Composition string `sheet:"Состав"`
		Description string `sheet:"Описание"`
		Research    string `sheet:"Исследования"`
		Matrix      string `sheet:"Рабочая матрица"`
		Prescribing string `sheet:"Рекомендации"`
	}

	SheetBanner struct {
		Image  string `sheet:"Изображение"`
		Active bool   `sheet:"Активен"`
		Href   string `sheet:"Переход"`
		Type   string `sheet:"Тип"`
	}

	SheetPrice struct {
		SKU   string `sheet:"Артикул"`
		Price int    `sheet:"Цена"`
	}

	SheetCDEKCity struct {
		Code         int      `sheet:"ID"`
		Name         string   `sheet:"CityName"`
		District     string   `sheet:"OblName"`
		PostCodeList []string `sheet:"PostCodeList"`
	}

	SheetRussiaPost struct {
		PostCode          string `sheet:""`
		Region            string `sheet:"Регион"`
		Capital           string `sheet:"Столица"`
		DeliveryTimeRapid string `sheet:"Время ускоренное"`
		DeliveryTimeEMC   string `sheet:"Время курьерское"`
	}

	SheetAccounting struct {
		ProductSKU string `sheet:"Артикул"`
		Amount   int    `sheet:"Количество"`
		Price      int    `sheet:"Цена"`
	}
)

func parseBannerType(label string) BannerType {
	switch label {
	case "main":
		return BannerTypeMain
	case "brand":
		return BannerTypeBrand
	}
	return ""
}

func CreatePositionReport(sheetData SheetAccounting) PositionReport {
	return PositionReport{
		ProductSKU: sheetData.ProductSKU,
		Amount: sheetData.Amount,
		Price: Price(sheetData.Price * 100),
	}
}

func CreateBanner(sheetData SheetBanner) Banner {
	return Banner{
		Image:  sheetData.Image,
		Href:   sheetData.Href,
		Active: sheetData.Active,
		Type:   parseBannerType(sheetData.Type),
	}
}

func CreateCDEKCity(sheetData SheetCDEKCity, postCode string) CDEKCity {
	return CDEKCity{
		Code:     sheetData.Code,
		Name:     sheetData.Name,
		District: sheetData.District,
		PostCode: postCode,
	}
}

func CreateNotation(sheetData SheetNotation) Notation {
	return Notation{
		SKU:         sheetData.SKU,
		Matrix:      sheetData.Matrix,
		Effects:     sheetData.Effects,
		Research:    sheetData.Research,
		Description: sheetData.Description,
		Prescribing: sheetData.Prescribing,
		Composition: sheetData.Composition,
	}
}

func CreateCollection(sheetData SheetCollection) Collection {
	return Collection{
		Name: sheetData.Name,
		SKU:  sheetData.SKU,
	}
}

/**
	ШxДxВ
 */
func CreateDimension(token string) Dimension {
	split := strings.Split(token, "x")

	if len(split) < 3 {
		return Dimension{}
	}

	width, err := strconv.Atoi(split[0])
	if err != nil {
		return Dimension{}
	}

	length, err := strconv.Atoi(split[1])
	if err != nil {
		return Dimension{}
	}

	height, err := strconv.Atoi(split[2])
	if err != nil {
		return Dimension{}
	}

	return Dimension{
		Width:  width,
		Height: height,
		Length: length,
	}
}

func CreateRussiaPostDeliveryPeriod(sheetData SheetRussiaPost) RussiaPostDeliveryPeriod {
	return RussiaPostDeliveryPeriod{
		Region:  sheetData.Region,
		Capital: sheetData.Capital,
		EMC:     tokenToDeliveryPeriod(sheetData.DeliveryTimeEMC),
		Rapid:   tokenToDeliveryPeriod(sheetData.DeliveryTimeRapid),
	}
}

func CreateProduct(sheetData SheetProduct) Product {
	product := Product{
		Name:          sheetData.Name,
		Producer:      sheetData.Producer,
		Form:          sheetData.Form,
		Factor:        sheetData.Factor,
		Weight:        sheetData.Weight,
		SKU:           sheetData.SKU,
		Quantity:      sheetData.Quantity,
		CollectionSKU: sheetData.CollectionSKU,
		Price:         Price(sheetData.Price * 100), // 100 копеек
		Pictures:      []string{},
		Dimension:     CreateDimension(sheetData.Dimension),
	}

	for _, p := range sheetData.Pictures {
		product.Pictures = append(product.Pictures, fmt.Sprintf("products/%s/%s.jpg", product.SKU, p))
	}

	if len(sheetData.Discount) > 0 {
		if num, percent, ok := percent(sheetData.Discount); ok && num > 0 {
			if percent {
				product.Discount = &Discount{
					Type:   DiscountTypePercentage,
					Amount: num,
				}
			} else {
				product.Discount = &Discount{
					Type:   DiscountTypeFixedAmount,
					Amount: num,
				}
			}
		}
	}

	return product
}
