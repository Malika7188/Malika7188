package goreloaded

import (
	"strings"
	"testing"
)
func TestLow(t *testing.T) {
	input := "i am BEAUTIFUL (low)"
	input2 := strings.Split(input, " ")
	Expeted := "i am beautiful"
	

	output := Low(input2)
	output2 := strings.Join(output, " ")

	if output2 != Expeted {
		t.Fatalf("expected: %v but ouput is : %v", Expeted,output2)
	}
}