package russiaPost

import (
	"net/http"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"github.com/moul/http2curl"
	"encoding/base64"
	"io"
	"github.com/pkg/errors"
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
	ErrAddressNotValid = errors.New("адрес неприемлем для доставки")
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

func (p *RussiaPost) doRequest(method, path string, body interface{}) (*http.Response, error) {
	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s/%s", HostName, Version, path)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))

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

func(p *RussiaPost) Backlog(request OrderRequest) (*OrderResponse, error) {
	resp, err := p.doRequest("PUT", "user/backlog", request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var response OrderResponse

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

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Ответ 404")
	}

	return &response, nil
}

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

func(p *RussiaPost) NormalizeName(request NormalizeNameRequest) (*NormalizeName, error) {
	resp, err := p.doRequest("POST", "clean/physical", []NormalizeNameRequest{request})

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Ответ 404")
	}

	var response []NormalizeName
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