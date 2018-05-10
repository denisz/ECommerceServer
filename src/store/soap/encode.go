package soap

import (
	"encoding/xml"
	"fmt"
)

// MarshalXML envelope the body and encode to xml
func (c Client) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {
	//start envelope
	if c.Definitions == nil {
		return fmt.Errorf("definitions is nil")
	}

	tokens, err := startToken(c.Definitions.TargetNamespace)
	if err != nil {
		return err
	}

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	wrap := wrapToken(c.Method)

	e.EncodeElement(c.Params, wrap)

	//e.Indent()

	//end envelope
	tokens = endToken()

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	return e.Flush()
}

/**
	xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
	xmlns:oper="http://russianpost.org/operationhistory"
	xmlns:data="http://russianpost.org/operationhistory/data"
	xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
 */
// startToken initiate body of the envelope
func startToken(n string) ([]xml.Token, error) {
	e := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Envelope",
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns:soap"}, Value: "http://www.w3.org/2003/05/soap-envelope"},
			{Name: xml.Name{Space: "", Local: "xmlns:oper"}, Value: n},
			{Name: xml.Name{Space: "", Local: "xmlns:data"}, Value: fmt.Sprintf("%s/data", n)},
			{Name: xml.Name{Space: "", Local: "xmlns:soapenv"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
		},
	}

	b := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Body",
		},
	}

	sh := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Header",
		},
	}

	eh := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Header",
		},
	}

	return []xml.Token{ e, sh, eh, b }, nil
}

// endToken close body of the envelope
func endToken() []xml.Token {
	e := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Envelope",
		},
	}

	b := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Body",
		},
	}

	return []xml.Token{b, e}
}

func wrapToken(m string) xml.StartElement {
	return xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: fmt.Sprintf("oper:%s", m),
		},
	}
}
