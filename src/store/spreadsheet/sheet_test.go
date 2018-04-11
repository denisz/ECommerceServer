package spreadsheet

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
	var data [][]interface{}

	s := make([]interface{}, 4)
	s[0] = "номер"
	s[1] = "опции"
	s[2] = "имя"
	s[3] = "картинка"
	data = append(data, s)

	s = make([]interface{}, 4)
	s[0] = "-1"
	s[1] = "1,2,3,4"
	s[2] = "v1"
	s[3] = "v2"
	data = append(data, s)

	s = make([]interface{}, 4)
	s[0] = "2"
	s[1] = "23,34,45"
	s[2] = "v1"
	s[3] = "v2"
	data = append(data, s)

	var sheet []MockSheet
	err := Unmarshal(data, &sheet)
	if err != nil {
		t.Error(err)
	}

	for _, i := range sheet {
		fmt.Printf("%v \n", i)
	}
}
