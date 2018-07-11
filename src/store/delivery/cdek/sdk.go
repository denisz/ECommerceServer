package cdek

import (
	"fmt"
	"time"
	"crypto/md5"
	"net/http"
	"bytes"
	"github.com/moul/http2curl"
	"io"
	"net/url"
	"encoding/xml"
	"io/ioutil"
)

const (
	SDKHostName = "https://integration.cdek.ru"
	SDKVersion = "1.0"
)

type SDK struct {
	Account string
	Password string
	Debug bool
}

var DefaultSDK = NewSDKClient(accountProd, passwordProd,true)
var DebugSDK = NewSDKClient(accountDev, passwordDev,true)

func NewSDKClient(Account string, Password string, Debug bool) *SDK {
	return &SDK{
		Account,
		Password,
		Debug,
	}
}

func (p *SDK) createUrl(path string) string {
	return fmt.Sprintf("%s/%s", SDKHostName, path)
}

func (p *SDK) createDateExecute() string {
	return time.Now().Format("2006-01-02")
}

func (p *SDK) createRequest(method, path string, body io.Reader) (*http.Request, error) {
	uri := p.createUrl(path)
	if method == "GET" {
		return http.NewRequest(method, uri, nil)
	}

	if method == "POST_EMPTY" {
		return http.NewRequest("POST", uri, nil)
	}

	return http.NewRequest(method, uri, body)
}

func (p *SDK) createSignature() string {
	data := []byte(fmt.Sprintf("%s&%s", p.createDateExecute(), p.Password))
	return fmt.Sprintf("%x", md5.Sum(data))
}

func (p *SDK) doRequest(method, path string, body interface{}) (*http.Response, error) {
	var bodyReader *bytes.Buffer = nil

	if body != nil {
		xmlEncode, err := xml.Marshal(body)
		if err != nil {
			return nil, err
		}
		xmlStr := fmt.Sprintf("%s%s", xml.Header, string(xmlEncode))

		q := url.Values{}
		q.Add("xml_request", xmlStr)

		bodyReader = bytes.NewBufferString(q.Encode())
	}

	req, err := p.createRequest(method, path, bodyReader)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if p.Debug {
		command, _ := http2curl.GetCurlCommand(req)
		fmt.Println(command)
	}

	client := &http.Client{}
	return client.Do(req)
}

//https://integration.cdek.ru/new_orders.php
func(p *SDK) CreateBacklog(request DeliveryRequest) (int, error) {
	resp, err := p.doRequest("POST", "new_orders.php", request)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	responseData,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	fmt.Printf("Response: %s \n", string(responseData))

	return 0, nil
}

//https://integration.cdek.ru/delete_orders.php.
func(p *SDK) DeleteBacklog(request OrderCreateRequest) (int, error) {
//delete_orders.php.
	return 0, nil
}

//https://integration.cdek.ru/orders_print.php.
func (p *SDK) Forms(id string, date time.Time) ([]byte, error) {
//orders_print.php.
	return []byte{}, nil
}
