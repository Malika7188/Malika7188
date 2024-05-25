package goreloaded

import (
	"strings"
	"testing"
)

func TestApostrophe(t *testing.T) {
	input := "' malika '"
	input2 := strings.Split(input, " ")
	Expeted := "'malika'"
	

	output := HandleApostrophes(input2)
	output2 := strings.Join(output, " ")

	if output2 != Expeted {
		t.Fatalf("expected: %v but ouput is : %v", Expeted,output2)
	}
}