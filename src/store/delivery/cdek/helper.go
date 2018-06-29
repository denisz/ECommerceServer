package cdek

import (
	"time"
	"fmt"
	"crypto/md5"
)

func CreateSecureSignature(password string, date time.Time) string {
	dateStr := date.Format("2006-01-02")
	data := []byte(fmt.Sprintf("%s&%s", dateStr, password))
	return fmt.Sprintf("%x", md5.Sum(data))
}