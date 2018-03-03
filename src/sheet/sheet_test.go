package sheet

import (
	"testing"
	"fmt"
)

type MockSheet struct {
	ID int `sheet:"номер"`
	Name string `sheet:"имя"`
	Options []string `sheet:"опции"`
	Picture string `sheet:"картинка"`
}

func TestMarshal(t *testing.T) {
	var sheet = []MockSheet{
		{
			Name: "test",
			Picture: "Test",
		},
	}
	data, err := Marshal(&sheet)

	if err != nil {
		t.Error(err)
	}
	fmt.Sprintf("%v", data)
}

func TestUnmarshal(t *testing.T) {
	data := [][]string {
		{"номер", "опции", "имя", "картинка"},
		{"-1", "1,2,3,4", "v1", "v2"},
		{"2", "23,34,45", "v1", "v2"},
	}

	var sheet []MockSheet
	err := Unmarshal(data, &sheet)
	if err != nil {
		t.Error(err)
	}

	for _, i := range sheet {
		fmt.Printf("%v \n", i)
	}
}
