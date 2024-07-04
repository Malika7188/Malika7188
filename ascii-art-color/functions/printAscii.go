package functions

import (
	"fmt"
	
)

// PrintChar takes a rune and font data, and returns the corresponding lines for that character
func PrintChar(c rune, fontData [95][8]string) ([]string, error) {
	var lines []string

	// Check if the rune is outside the printable ASCII range (32-126)
	if c > 126 || c < 32 {
		return nil, fmt.Errorf("gross error reading file")
	}

	// Calculate the index of the character in the fontData array
	charIndex := int(c) - 32

	// Append each line of the character's font data to the lines slice
	for i := 0; i < 8; i++ {
		lines = append(lines, fontData[charIndex][i])
	} 
	

	// Return the lines slice and a nil error to indicate success
	return lines, nil
}
