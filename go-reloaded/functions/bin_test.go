package goreloaded

import (
	"strings"
	"testing"
)

func TestBin(t *testing.T) {
	input := "10 (bin) years"
	input2 := strings.Split(input, " ")
	Expeted := "2 years"

	output := Bin(input2)
	output2 := strings.Join(output, " ")

	if output2 != Expeted {
		t.Fatalf("expected: %v but ouput is : %v", Expeted, output2)
	}
}
