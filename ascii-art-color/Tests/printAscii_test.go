package tests

import (
	"os"
	"strings"
	"testing"

	"ascii-art-color/functions"
)

func TestPrintChar(t *testing.T) {
	// Load the font data
	fontArray, err := functions.FontLoader("standard")
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

	outputs, _ := functions.PrintChar('L', fontArray)

	// Trim trailing whitespace and newline characters from actual output
	actual := strings.TrimSpace(strings.Join(outputs, "\n"))

	// Compare expected output with actual output
	if actual != expected {
		t.Errorf("expected: %v , but got: %v", expected, actual)
	}
}
