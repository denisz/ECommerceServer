package cdek

import (
	"fmt"
	"time"
	"crypto/md5"
)

const (
	SDKHostName = "https://integration.cdek.ru/"
	SDKVersion = "1.0"
)

type SDK struct {
	Account string
	Password string
	Debug bool
}

func (p *SDK) createUrl(path string) string {
	return fmt.Sprintf("%s/%s", HostName, path)
}

func (p *SDK) createDateExecute() string {
	return time.Now().Format("2006-01-02")
}

func (p *SDK) createSignature() string {
	data := []byte(fmt.Sprintf("%s&%s", p.createDateExecute(), p.Password))
	return fmt.Sprintf("%x", md5.Sum(data))
}

func(p *SDK) CreateBacklog(request OrderCreateRequest) (int, error) {
//new_orders.php;

	return 0, nil
}

func(p *SDK) DeleteBacklog(request OrderCreateRequest) (int, error) {
//delete_orders.php.
	return 0, nil
}

func (p *SDK) Forms(id string, date time.Time) ([]byte, error) {
//orders_print.php.
	return []byte{}, nil
}
