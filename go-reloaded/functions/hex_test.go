package goreloaded

import (
	"strings"
	"testing"
)

func TestHex(t *testing.T) {
	input := "42 (hex) years"
	input2 := strings.Split(input, " ")
	Expeted := "66 years"

	output := Hex(input2)
	output2 := strings.Join(output, " ")

	if output2 != Expeted {
		t.Fatalf("expected: %v but ouput is : %v", Expeted, output2)
	}
}
