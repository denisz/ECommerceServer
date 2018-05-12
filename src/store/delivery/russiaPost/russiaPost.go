package russiaPost

import (
	"net/http"
	"bytes"
	"fmt"
	"encoding/json"
	"github.com/moul/http2curl"
	"encoding/base64"
	"io"
	"github.com/pkg/errors"
	"strings"
	"time"
	"io/ioutil"
)

const (
	HostName = "https://otpravka-api.pochta.ru"
	Version = "1.0"
)

type RussiaPost struct {
	Login string
	Password string
	AccessToken string
	Debug bool
}

var (
	ErrAddressNotValid = errors.New("Адрес неприемлем для доставки")
	ErrPhysicalNotValid = errors.New("Имя отрпавителя неприемлем для доставки")
	ErrPhoneNotValid = errors.New("Номер телефона неприемлем для доставки")
)

const (
	token = "MmmDeJqGxRlL2MXX4oZiknt25K5mUFEg"
	login = "denisxy12@hotmail.com"
	password = "2Q2sminvc"
)

var DefaultClient = NewClient(login, password, token, true)

func NewClient(Login string, Password string, AccessToken string, Debug bool) *RussiaPost {
	return &RussiaPost{
		Login,
		Password,
		AccessToken,
		Debug,
	}
}

func (p *RussiaPost) createUrl(path string) string {
	return fmt.Sprintf("%s/%s/%s", HostName, Version, path)
}

func (p *RussiaPost) createRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := p.createUrl(path)
	if method == "GET" {
		return http.NewRequest(method, url, nil)
	}
	return http.NewRequest(method, url, body)
}

func (p *RussiaPost) doRequest(method, path string, body interface{}) (*http.Response, error) {
	var bodyReader *bytes.Buffer = nil

	if body != nil {
		jsonStr, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		bodyReader = bytes.NewBuffer(jsonStr)
	}

	req, err := p.createRequest(method, path, bodyReader)

	if err != nil {
		return nil, err
	}

	loginPassword := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", p.Login, p.Password)))
	req.Header.Set("Authorization", fmt.Sprintf("AccessToken %s", p.AccessToken))
	req.Header.Set("X-User-Authorization", fmt.Sprintf("Basic %s", loginPassword))
	req.Header.Set("Accept", "application/json;charset=UTF-8")
	req.Header.Set("Content-Type", "application/json")

	if p.Debug {
		command, _ := http2curl.GetCurlCommand(req)
		fmt.Println(command)
	}

	client := &http.Client{}
	return client.Do(req)
}

//Создает новый заказ. Автоматически рассчитывает и проставляет плату за пересылку.
func(p *RussiaPost) CreateBacklog(request OrderRequest) (int, error) {
	resp, err := p.doRequest("PUT", "user/backlog", []OrderRequest{request})

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	var response CreateEntityResponse

	dec := json.NewDecoder(resp.Body)

	for {
		if err := dec.Decode(&response); err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}
	}

	if p.Debug {
		strJson, _ := json.Marshal(response)
		fmt.Printf("Response: %v %v", string(strJson), resp.StatusCode)
	}

	if len(response.Errors) > 0 {
		errDescription := strings.Builder{}
		codes := response.Errors[0].Codes
		for _, code := range codes {
			errDescription.WriteString(code.Details)
			errDescription.WriteString(". ")
		}
		return 0, errors.New(errDescription.String())
	}

	return response.Ids[0], nil
}

//Удаление заказа
func (p *RussiaPost) DeleteBacklog(id string) (*CreateEntityResponse, error) {
	resp, err := p.doRequest("DELETE", "user/backlog", []string{id})

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var response CreateEntityResponse

	dec := json.NewDecoder(resp.Body)

	for {
		if err := dec.Decode(&response); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	if p.Debug {
		strJson, _ := json.Marshal(response)
		fmt.Printf("Response: %v %v", string(strJson), resp.StatusCode)
	}

	if len(response.Errors) > 0 {
		errDescription := strings.Builder{}
		codes := response.Errors[0].Codes
		for _, code := range codes {
			errDescription.WriteString(code.Details)
		}
		return nil, errors.New(errDescription.String())
	}

	return &response, nil
}

//Метод переводит заказы из партии в раздел Новые. Партия должна быть в статусе CREATED.
func (p *RussiaPost) RestoreBacklog(ids []string) (*CreateEntityResponse, error) {
	resp, err := p.doRequest("POST", "user/backlog", ids)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var response CreateEntityResponse

	dec := json.NewDecoder(resp.Body)

	for {
		if err := dec.Decode(&response); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	if p.Debug {
		strJson, _ := json.Marshal(response)
		fmt.Printf("Response: %v %v", string(strJson), resp.StatusCode)
	}

	if len(response.Errors) > 0 {
		errDescription := strings.Builder{}
		codes := response.Errors[0].Codes
		for _, code := range codes {
			errDescription.WriteString(code.Details)
		}
		return nil, errors.New(errDescription.String())
	}

	return &response, nil
}

// Генерирует и возвращает pdf файл, который содержит форму Е-1
// (EMS, EMS-оптимальное, Бизнес курьер, Бизнес курьер экспресс) для указанного заказа
func (p *RussiaPost) FormsE1(id string, date time.Time) ([]byte, error) {
	path := fmt.Sprintf("forms/backlog/%s/forms?sending-date=%s", id, date.Format("2006-01-02"))
	resp, err := p.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Not found")
	}
	return ioutil.ReadAll(resp.Body)
}

// Генерирует и возвращает pdf файл с формой ф7п для указанного заказа.
func (p *RussiaPost) FormsF7(id string, date time.Time) ([]byte, error) {
	path := fmt.Sprintf("forms/%s/f7pdf?sending-date=%s&print-type=PAPER", id, date.Format("2006-01-02"))
	resp, err := p.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Not found")
	}
	return ioutil.ReadAll(resp.Body)
}

// Генерирует и возвращает pdf файл с формой Ф103 для указанной партии.
func (p *RussiaPost) FormsF103(name string) ([]byte, error) {
	path := fmt.Sprintf("forms/%s/f103pdf", name)
	resp, err := p.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Not found")
	}
	return ioutil.ReadAll(resp.Body)
}

/** Генерирует и возвращает zip архив с 4-мя файлами:
	Export.xls , Export.csv - список с основными данными по заявкам в составе партии
	F103.pdf - форма ф103 по заявкам в составе партии
	В зависимости от типа и категории отправлений, формируется комбинация из
	сопроводительных документов в формате pdf ( формы: f7, f112, f22)
 */
func (p *RussiaPost) FormsPacket(name string) ([]byte, error) {
	path := fmt.Sprintf("forms/%s/zip-all", name)
	resp, err := p.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Not found")
	}
	return ioutil.ReadAll(resp.Body)
}

//Поиск заказа по идентификатору
func (p *RussiaPost) GetBacklog(id int) (*Order, error) {
	path := fmt.Sprintf("backlog/%d", id)
	resp, err := p.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Not found")
	}

	var response Order
	dec := json.NewDecoder(resp.Body)
	for {
		if err := dec.Decode(&response); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	if p.Debug {
		strJson, _ := json.Marshal(response)
		fmt.Printf("Response: %v %v", string(strJson), resp.StatusCode)
	}

	return &response, nil
}

/**
	Автоматически создает партию и переносит указанные подготовленные заказы в эту партию.
	Если заказы относятся к разным типам и категориям – создается несколько партий.
	Заказы распределяются по соответствующим партиям. Каждому перенесенному заказу автоматически присваивается ШПИ.
 */
func (p *RussiaPost) Shipment(ids []string, date time.Time) (*BatchesResponse, error) {
	path := fmt.Sprintf("user/shipment?sending-date=%s", date.Format("2006-01-02"))
	resp, err := p.doRequest("POST", path, ids)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response BatchesResponse
	dec := json.NewDecoder(resp.Body)

	for {
		if err := dec.Decode(&response); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	if p.Debug {
		strJson, _ := json.Marshal(response)
		fmt.Printf("Response: %v %v", string(strJson), resp.StatusCode)
	}

	if len(response.Errors) > 0 {
		errDescription := strings.Builder{}
		codes := response.Errors[0].Codes
		for _, code := range codes {
			errDescription.WriteString(code.Details)
		}
		return nil, errors.New(errDescription.String())
	}

	return &response, nil
}

//Рассчитывает стоимость пересылки в зависимости от указанных входных данных.
func(p *RussiaPost) Tariff(request DestinationRequest) (*DestinationResponse, error) {
	resp, err := p.doRequest("POST", "tariff", request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respError DestinationError
		dec := json.NewDecoder(resp.Body)

		for {
			if err := dec.Decode(&respError); err == io.EOF {
				break
			} else if err != nil {
				return nil, err
			}
		}

		if p.Debug {
			strJson, _ := json.Marshal(respError)
			fmt.Printf("error: %v %v \n", string(strJson), resp.StatusCode)
		}
		return nil, errors.New(respError.DescError)
	}

	var response DestinationResponse
	dec := json.NewDecoder(resp.Body)

	for {
		if err := dec.Decode(&response); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	if p.Debug {
		strJson, _ := json.Marshal(response)
		fmt.Printf("response: %v \n", string(strJson))
	}

	return &response, nil
}

//Нормализация адреса
func(p *RussiaPost) NormalizeAddress(request NormalizeAddressRequest) (*NormalizeAddress, error) {
	resp, err := p.doRequest("POST", "clean/address", []NormalizeAddressRequest{request})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Ответ 404")
	}

	var response []NormalizeAddress
	dec := json.NewDecoder(resp.Body)

	for {
		if err := dec.Decode(&response); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	if p.Debug {
		strJson, _ := json.Marshal(&response[0])
		fmt.Printf("Response: %v \n", string(strJson))
	}

	return &response[0], nil
}

//Нормализация ФИО
func(p *RussiaPost) NormalizePhysical(request NormalizePhysicalRequest) (*NormalizePhysical, error) {
	resp, err := p.doRequest("POST", "clean/physical", []NormalizePhysicalRequest{request})

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Ответ 404")
	}

	var response []NormalizePhysical
	dec := json.NewDecoder(resp.Body)

	for {
		if err := dec.Decode(&response); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	if p.Debug {
		strJson, _ := json.Marshal(&response[0])
		fmt.Printf("Response: %v \n", string(strJson))
	}

	return &response[0], nil
}

//Нормализация телефона
func (p *RussiaPost) NormalizePhone(request NormalizePhoneRequest) (*NormalizePhone, error) {
	resp, err := p.doRequest("POST", "clean/phone", []NormalizePhoneRequest{request})

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Ответ 404")
	}

	var response []NormalizePhone
	dec := json.NewDecoder(resp.Body)

	for {
		if err := dec.Decode(&response); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	if p.Debug {
		strJson, _ := json.Marshal(&response[0])
		fmt.Printf("Response: %v \n", string(strJson))
	}

	return &response[0], nil
}