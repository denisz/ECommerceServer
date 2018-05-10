package updater

import (
	"store/spreadsheet"
	"regexp"
	"strconv"
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

func percent(token string) (float64, bool, bool) {
	r, _ := regexp.Compile(`^([0-9.]+)([%]?)$`)
	t := r.FindStringSubmatch(token)

	if len(t) == 3 {
		if t[2] == "%" {
			i, err := strconv.Atoi(t[1])
			if err != nil {
				return 0, false, false
			}

			return float64(i), true, true
		}

		i, err := strconv.Atoi(t[1])
		if err != nil {
			return 0, false, false
		}
		return float64(i), false, true
	}

	return 0, false, false
}