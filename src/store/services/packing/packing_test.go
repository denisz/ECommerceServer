package packing

import (
	"testing"
	"fmt"
)

func TestPacker(t *testing.T) {
	p := NewPacker()

	// Add bins.
	p.AddBin(NewBin("Small Bin", 10, 15, 20, 100))
	p.AddBin(NewBin("Medium Bin", 100, 150, 200, 1000))

	// Add items.
	p.AddItem(NewItem("Item 1", 2, 2, 1, 200))
	p.AddItem(NewItem("Item 2", 3, 3, 2, 300))

	// Pack items to bins.
	if err := p.Pack(); err != nil {
		t.Fatal(err)
	}

	for _, b := range p.Bins {
		fmt.Println(b)
		fmt.Println(" packed items:")
		for _, i := range b.Items {
			fmt.Println("  ", i)
		}
	}

	// Pack items to bins.
	if err := p.Pack(); err != nil {
		t.Fatal(err)
	}
}