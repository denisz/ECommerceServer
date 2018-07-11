package cdek

import (
	"fmt"
	"crypto/md5"
)

func CreateSecureSignature(password string, date string) string {
	data := []byte(fmt.Sprintf("%s&%s", date, password))
	return fmt.Sprintf("%x", md5.Sum(data))
}