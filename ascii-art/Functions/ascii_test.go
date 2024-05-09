package functions

import (
	//"strings"
	"fmt"
	"strings"
	"testing"
)

func TestAscciChar(t *testing.T) {
	// Load the font data
	fontArray, err := Fontloader("standard.txt")
	if err != nil {
		fmt.Println("Error loading font:", err)
		return
	}

	outputs, _ := AsciiChar('H', fontArray)
	slice := []string{
		"   _    _  ",
		"  | |  | |  ",
		"  | |__| |  ",
		"  |  __  |  ",
		"  | |  | |  ",
		"  |_|  |_|  ",
	}
	output := slice
	if strings.Join(outputs, "\n") != strings.Join(output, "\n") {
		t.Errorf("expected: %v , but got: %v", strings.Join(slice, "\n"), string("\n"+strings.Join(outputs, "\n")))
	}

	// fmt.Println(expected)
}
