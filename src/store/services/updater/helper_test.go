package updater

import (
	"testing"
	"fmt"
)

func TestUnderscoreString(t *testing.T) {
	str := "FG RT 50"
	result := underscoreString(str)

	fmt.Print(result)
	if result != "FG_RT_50" {
		t.Error("Expected correct result")
	}

	str = "fg RT 50"
	result = underscoreString(str)

	fmt.Print(result)
	if result != "fg_RT_50" {
		t.Error("Expected correct result")
	}
}
