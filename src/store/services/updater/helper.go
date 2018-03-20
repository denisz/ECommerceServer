package updater

import (
	"fmt"
	"spreadsheet"
	"regexp"
	"strconv"
	"strings"
)

func TransformToStrings(data [][]interface{}) [][]string {
	strings := make([][]string, len(data))
	for i, row := range data {
		strings[i] = make([]string, len(row))
		for j, cell := range row {
			strings[i][j] = fmt.Sprint(cell)
		}
	}

	return strings
}

func UnmarshalSpreadsheet(out interface{}, spreadsheetId string, readRange string) (err error) {
	data, err := ReadSpreadsheet(spreadsheetId, readRange)
	if err != nil {
		return
	}

	err = spreadsheet.Unmarshal(TransformToStrings(data), out)
	if err != nil {
		return
	}

	return
}

func percent(token string) (int, bool, bool) {
	r, _ := regexp.Compile(`^([0-9]+)([%]?)$`)
	t := r.FindStringSubmatch(token)

	if len(t) == 3 {
		if t[2] == "%" {
			i, err := strconv.Atoi(t[1])
			if err != nil {
				return 0, false, false
			}

			return i, true, true
		}

		i, err := strconv.Atoi(t[1])
		if err != nil {
			return 0, false, false
		}
		return i, false, true
	}

	return 0, false, false
}

func underscoreString(str string) string {

	// convert every letter to lower case
	newStr := str //strings.ToLower(str)

	// convert all spaces/tab to underscore
	regExp := regexp.MustCompile("[[:space:][:blank:]]")
	newStrByte := regExp.ReplaceAll([]byte(newStr), []byte("_"))

	regExp = regexp.MustCompile("`[^a-z0-9]`i")
	newStrByte = regExp.ReplaceAll(newStrByte, []byte("_"))

	regExp = regexp.MustCompile("[!/']")
	newStrByte = regExp.ReplaceAll(newStrByte, []byte("_"))

	// and remove underscore from beginning and ending

	newStr = strings.TrimPrefix(string(newStrByte), "_")
	newStr = strings.TrimSuffix(newStr, "_")

	return newStr
}
