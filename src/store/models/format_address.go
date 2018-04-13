package models

import (
	"strings"
)

func (p *Address) Format() string {
	if p.ManualInput {
		return strings.Join([]string {
			p.PostalCode,
			p.Country,
			p.Region,
			p.City,
			p.Street,
			p.House,
			p.Room,
			p.Comment,
		}, ", ")
	}

	return strings.Join([]string{
		p.PostalCode,
		p.Address,
		p.Comment,
	}, ", ")
}
