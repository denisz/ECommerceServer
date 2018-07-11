package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
	. "store/models"
	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/gob"
)

func TestCreateInvoice(t *testing.T) {
	invoice1, _ := CreateInvoice()
	invoice2, _ := CreateInvoice()

	fmt.Printf("invoice1: %v \n", invoice1)
	fmt.Printf("invoice2: %v \n", invoice2)
	assert.NotEqual(t, invoice1, invoice2)
}

func TestControllerCart_GetDeliveryPeriodForRussiaPost(t *testing.T) {
	var periods []RussiaPostDeliveryPeriod
	DB, err := storm.Open("../../../../db/store.db", storm.Codec(gob.Codec))
	if err != nil {
		t.Error(err)
	}

	err = DB.From(NodeNamedRussiaPost).AllByIndex("ID", &periods)

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v", periods)
}