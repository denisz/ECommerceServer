package soap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Params type is used to set the params in soap request
type Params interface{}

// SoapClient return new *Client to handle the requests with the WSDL
func SoapClient(wsdl string, debug bool) (*Client, error) {
	_, err := url.Parse(wsdl)
	if err != nil {
		return nil, err
	}

	d, err := getWsdlDefinitions(wsdl)
	if err != nil {
		return nil, err
	}

	c := &Client{
		WSDL:        wsdl,
		URL:         strings.TrimSuffix(d.TargetNamespace, "/"),
		Definitions: d,
		Debug:       debug,
	}

	return c, nil
}

// Client struct hold all the informations about WSDL,
// request and response of the server
type Client struct {
	WSDL        string
	URL         string
	Method      string
	Params      Params
	Definitions *wsdlDefinitions
	Body        []byte
	Debug       bool
	payload     []byte
}

// Call call's the method m with Params p
func (c *Client) Call(m string, p Params) (err error) {
	c.Method = m
	c.Params = p

	xmlstring, err := xml.Marshal(c)
	xmlstring = []byte(xml.Header + string(xmlstring))
	c.payload = xmlstring
	if c.Debug {
		fmt.Printf("Payload %s \n", string(c.payload))
	}

	if err != nil {
		return err
	}

	b, err := c.doRequest()
	if c.Debug {
		fmt.Printf("Response %s", string(b))
	}

	if err != nil {
		return err
	}

	var soap SoapEnvelope
	err = xml.Unmarshal(b, &soap)

	c.Body = soap.Body.Contents
	fmt.Printf("Body %s", string(c.Body))

	return err
}

// Unmarshal get the body and unmarshal into the interface
func (c *Client) Unmarshal(v interface{}) error {
	if len(c.Body) == 0 {
		return fmt.Errorf("Body is empty")
	}

	var f Fault
	xml.Unmarshal(c.Body, &f)
	if f.Code != "" {
		return fmt.Errorf("[%s]: %s", f.Code, f.Description)
	}

	return xml.Unmarshal(c.Body, v)
}

// doRequest makes new request to the server using the c.Method, c.URL and the body.
// body is enveloped in Call method
func (c *Client) doRequest() ([]byte, error) {
	req, err := http.NewRequest("POST", c.WSDL, bytes.NewBuffer(c.payload))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req.ContentLength = int64(len(c.payload))

	req.Header.Add("Content-Type", "application/soap+xml; charset=utf-8")
	req.Header.Add("Accept", "text/xml")
	req.Header.Add("SOAPAction", "")
	req.Header.Set("User-Agent", "golang")

	req.Close = true

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// SoapEnvelope struct
type SoapEnvelope struct {
	XMLName struct{} `xml:"Envelope"`
	Body    SoapBody
}

// SoapBody struct
type SoapBody struct {
	XMLName  struct{} `xml:"Body"`
	Contents []byte   `xml:",innerxml"`
}
