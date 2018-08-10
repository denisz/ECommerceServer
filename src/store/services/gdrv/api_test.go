package gdrv

import (
	"testing"
	"fmt"
)

func TestWalk(t *testing.T) {
	files, err := Walk("1br5HTFWzYWTrALqfBb3SYWbFZOzhe8lj")

	if err != nil {
		t.Error(err)
	}

	for _, file := range files {
		cells, err := ReadSpreadsheet(file.Id, "Sheet1")
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("%v \n", cells)
	}
}


func TestGetAllSheets(t *testing.T) {
	files, err := Walk("1br5HTFWzYWTrALqfBb3SYWbFZOzhe8lj")

	if err != nil {
		t.Error(err)
	}

	for _, file := range files {
		sheets, err := GetAllSheets(file.Id)
		if err != nil {
			t.Error(err)
		}

		for _, sh := range sheets {
			fmt.Printf("%v \n", sh.Properties.Title)
		}
	}
}