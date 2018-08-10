package gdrv

import (
	"store/spreadsheet"
)

func UnmarshalSpreadsheet(out interface{}, spreadsheetId string, readRange string) (err error) {
	data, err := ReadSpreadsheet(spreadsheetId, readRange)
	if err != nil {
		return
	}

	err = spreadsheet.Unmarshal(data, out)
	if err != nil {
		return
	}

	return
}