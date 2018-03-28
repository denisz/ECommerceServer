package pochta

import (
	"net/http"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"github.com/moul/http2curl"
	"encoding/base64"
	"io"
	"io/ioutil"
	"github.com/pkg/errors"
)

const (
	HostName = "https://otpravka-api.pochta.ru"
	Version = "1.0"
)

type Pochta struct {
	Login string
	Password string
	AccessToken string
	Debug bool
}

func NewPochta(Login string, Password string, AccessToken string, Debug bool) *Pochta {
	return &Pochta{
		Login,
		Password,
		AccessToken,
		Debug,
	}
}

func(p *Pochta) Backlog(request *OrderRequest) (*OrderResponse, error) {
	jsonStr, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s/user/backlog", HostName, Version)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	loginPassword := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", p.Login, p.Password)))
	req.Header.Set("Authorization", fmt.Sprintf("AccessToken %s", p.AccessToken))
	req.Header.Set("X-User-Authorization", fmt.Sprintf("Basic %s", loginPassword))
	req.Header.Set("Accept", "application/json;charset=UTF-8")
	req.Header.Set("Content-Type", "application/json")

	command, _ := http2curl.GetCurlCommand(req)
	fmt.Println(command)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		responseData, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("response: %v %v \n", string(responseData), resp.StatusCode)

		return nil, errors.New("Ответ 404")
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

	fmt.Printf("Response: %v", response)

	return &response, nil
}

func(p *Pochta) Tariff(request *DestinationRequest) (*DestinationResponse, error) {
	jsonStr, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s/tariff", HostName, Version)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	loginPassword := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", p.Login, p.Password)))
	req.Header.Set("Authorization", fmt.Sprintf("AccessToken %s", p.AccessToken))
	req.Header.Set("X-User-Authorization", fmt.Sprintf("Basic %s", loginPassword))
	req.Header.Set("Accept", "application/json;charset=UTF-8")
	req.Header.Set("Content-Type", "application/json")

	command, _ := http2curl.GetCurlCommand(req)
	fmt.Println(command)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		responseData, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("response: %v %v \n", string(responseData), resp.StatusCode)
		return nil, errors.New("Ответ 404")
	}

	defer resp.Body.Close()
	var response DestinationResponse

	dec := json.NewDecoder(resp.Body)

	for {
		if err := dec.Decode(&response); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	return &response, nil
}
