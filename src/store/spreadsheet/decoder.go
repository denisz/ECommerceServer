package spreadsheet

import "strings"

type Decoder struct {
	Rows chan *Row
}

type Row struct {
	Index int
	Cells chan *Cell
}

type Cell struct {
	Header string
	Token string
}

func NewDecoder(in [][]string) *Decoder {
	decoder := Decoder{
		make(chan *Row),
	}

	go func() {
		var headers []string

		if len(in) == 0 {
			close(decoder.Rows)
			return
		}

		for _, cell := range in[0] {
			headers = append(headers, cell)
		}

		for i := 1; i < len(in); i++ {
			row := Row{
				Index: i - 1,
				Cells: make(chan *Cell),
			}

			decoder.Rows <- &row

			for idx := 0; idx < len(in[i]); idx++ {
				token := in[i][idx]
				token = strings.TrimSpace(token)

				if len(headers) <= idx || len(token) == 0 {
					continue
				}

				header := headers[idx]

				row.Cells <- &Cell { header, token }
			}

			close(row.Cells)
		}

		close(decoder.Rows)
	}()

	return &decoder
}
