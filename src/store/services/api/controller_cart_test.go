package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func TestCreateInvoice(t *testing.T) {
	invoice1, _ := CreateInvoice()
	invoice2, _ := CreateInvoice()

	fmt.Printf("invoice1: %v \n", invoice1)
	fmt.Printf("invoice2: %v \n", invoice2)
	assert.NotEqual(t, invoice1, invoice2)
}