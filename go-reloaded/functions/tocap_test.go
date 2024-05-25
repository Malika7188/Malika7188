package goreloaded

import (
	"strings"
	"testing"
)

func TestCap(t *testing.T) {
	input := "i (cap) am beautiful"
	input2 := strings.Split(input, " ")
	Expected := "I am beautiful"

	output := Cap(input2)
	output2 := strings.Join(output, " ")

	if output2 != Expected {
		t.Fatalf("expected: %v but ouput is : %v", Expected, output2)
	}
}
