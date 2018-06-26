package cdek

import (
	"net/http"
	"bytes"
	"fmt"
	"github.com/moul/http2curl"
	"encoding/json"
	"io"
	"crypto/md5"
	"time"
	"errors"
)

const (
	HostName = "https://api.cdek.ru"
	HostNameIntegration = "https://integration.cdek.ru/"
	Version = "1.0"
)


type CDEK struct {
	Account string
	Password string
	Debug bool
	SDK SDK
}

var (
	ErrAddressNotValid = errors.New("Адрес неприемлем для доставки")
	ErrPhysicalNotValid = errors.New("Имя отрпавителя неприемлем для доставки")
	ErrPhoneNotValid = errors.New("Номер телефона неприемлем для доставки")
)

const (
	accountDev = "b2081ded55f6f254ec3eb9e32513ee9c"
	passwordDev = "ffa93d2a4ed42c8c6e39d33fd8820cb9"
	accountProd = "63d961a3f2c2a46af8f5c4dbd9ca205b"
	passwordProd = "631ed0d46ab1ece9544ef40346f2b971"
)

var DefaultClient = NewClient(accountProd, passwordProd,true)
var DebugClient = NewClient(accountDev, passwordDev,true)

func NewClient(Account string, Password string, Debug bool) *CDEK {
	return &CDEK{
		Account,
		Password,
		Debug,
		SDK{
			Account,
			Password,
			Debug,
		},
	}
}

func (p *CDEK) createUrl(path string) string {
	return fmt.Sprintf("%s/%s", HostName, path)
}

func (p *CDEK) createDateExecute() string {
	return time.Now().Format("2006-01-02")
}

func (p *CDEK) createSignature() string {
	data := []byte(fmt.Sprintf("%s&%s", p.createDateExecute(), p.Password))
	return fmt.Sprintf("%x", md5.Sum(data))
}

func (p *CDEK) createRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := p.createUrl(path)
	if method == "GET" {
		return http.NewRequest(method, url, nil)
	}

	if method == "POST_EMPTY" {
		return http.NewRequest("POST", url, nil)
	}

	return http.NewRequest(method, url, body)
}

func (p *CDEK) doRequest(method, path string, body interface{}) (*http.Response, error) {
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

	//req.Header.Set("Accept", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json;charset=UTF-8")
	req.Header.Set("Content-Type", "application/json")

	if p.Debug {
		command, _ := http2curl.GetCurlCommand(req)
		fmt.Println(command)
	}

	client := &http.Client{}
	return client.Do(req)
}


func(p *CDEK) Tariff(request DestinationRequest) (*DestinationResponse, error) {
	request.Version = Version
	request.Account = p.Account
	request.Secure = p.createSignature()
	request.DateExecute = p.createDateExecute()

	resp, err := p.doRequest("POST", "calculator/calculate_price_by_json.php", request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var response SuccessResponse

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
		fmt.Printf("Response: %v %v \n", string(strJson), resp.StatusCode)
	}

	return &response.Result, nil
}