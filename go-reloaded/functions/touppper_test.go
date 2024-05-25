package goreloaded

import (
	"strings"
	"testing"
)

func TestUp(t *testing.T) {
	input := "I am beautiful (up) and wonderful (up, 2)"
	input2 := strings.Split(input, " ")
	Expeted := "I am BEAUTIFUL AND WONDERFUL"

	output := Up(input2)
	output2 := strings.Join(output, " ")

	if output2 != Expeted {
		t.Fatalf("expected: %v but ouput is : %v", Expeted, output2)
	}
}
