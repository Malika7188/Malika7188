package tests

import (
	"os"
	"strings"
	"testing"

	functions "ascii-art/Functions"
)

func TestAscciChar(t *testing.T) {
	// Load the font data
	fontArray, err := functions.Fontloader("standard.txt")
	if err != nil {
		t.Fatalf("Error loading font: %v", err)
	}

	// Read ASCII character from file
	charContent, err := os.ReadFile("test.txt")
	if err != nil {
		t.Fatalf("Error reading test.txt: %v", err)
	}

	// Trim trailing whitespace and newline characters
	expected := strings.TrimSpace(string(charContent))

	outputs, _ := functions.AsciiChar('L', fontArray)

	// Trim trailing whitespace and newline characters from actual output
	actual := strings.TrimSpace(strings.Join(outputs, "\n"))

	// Compare expected output with actual output
	if actual != expected {
		t.Errorf("expected: %v , but got: %v", expected, actual)
	}
}
