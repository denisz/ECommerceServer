package spreadsheet


import (
	"reflect"
	"github.com/fatih/structtag"
	"errors"
	"strconv"
	"strings"
	"fmt"
)

var (
	ErrOutIsNotSlice = errors.New("out kind is not Slice")
	ErrNotImpl = errors.New("function is not implementation")
)

type Header struct {
	Name string
	Field string
	Options []string
}

func Marshal(in interface{}) (out interface{}, err error) {
	return nil, ErrNotImpl
}

func Unmarshal(data [][]interface{}, out interface{}) (err error) {

	in := make([][]string, len(data))
	for i, row := range data {
		in[i] = make([]string, len(row))
		for j, cell := range row {
			in[i][j] = fmt.Sprint(cell)
		}
	}

	ptyp := reflect.TypeOf(out)
	pval := reflect.ValueOf(out)

	var typ reflect.Type
	var val reflect.Value

	if ptyp.Kind() == reflect.Ptr {
		typ = ptyp.Elem()
		val = pval.Elem()
	} else {
		typ = ptyp
		val = pval
	}

	if val.Kind() != reflect.Slice {
		return ErrOutIsNotSlice
	}

	elemType := typ.Elem()

	var fields []Header

	for i := 0; i != elemType.NumField(); i++ {
		field := elemType.Field(i)
		if field.PkgPath != "" && !field.Anonymous {
			continue // Private field
		}
		// get field tag
		tag := field.Tag
		// ... and start using structtag by parsing the tag
		tags, err := structtag.Parse(string(tag))
		if err != nil {
			panic(err)
		}

		// iterate over all tags
		for _, t := range tags.Tags() {
			if t.Key == "sheet" {
				fields = append(fields, Header {
					Name: t.Name,
					Field: field.Name,
					Options: t.Options,
				})
			}
		}
	}

	decoder := NewDecoder(in)
	for row := range decoder.Rows {
		newItem := reflect.Indirect(reflect.New(elemType))
		for cell := range row.Cells {
			for _, field := range fields {
				if field.Name == cell.Header {
					f := newItem.FieldByName(field.Field)
					if f.CanSet() {
						kind := f.Kind()
						switch kind {
						case reflect.String:
							f.SetString(cell.Token)
						case reflect.Int:
							if num, err := strconv.ParseInt(cell.Token, 10, 0); err == nil {
								f.SetInt(num)
							}
						case reflect.Float64:
							if num, err := strconv.ParseFloat(cell.Token, 0); err == nil {
								f.SetFloat(num)
							}
						case reflect.Float32:
							if num, err := strconv.ParseFloat(cell.Token, 0); err == nil {
								f.SetFloat(num)
							}
						case reflect.Bool:
							if b, err := strconv.ParseBool(cell.Token); err == nil {
								f.SetBool(b)
							}
						case reflect.Slice:
							var slice []string
							candidates := strings.Split(cell.Token, ",")
							for _, s := range candidates {
								slice = append(slice, strings.Trim(s, " "))
							}

							f.Set(reflect.ValueOf(slice))
						}
					}
				}
			}
		}
		val.Set(reflect.Append(val, newItem))
	}

	return nil
}






