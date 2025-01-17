package functions

import "fmt"

func AsciiChar(C rune, fontData [95][8]string) ([]string, error) {
	// Check if the character is ASCII
	if C < 32 || C > 126 {
		return nil, fmt.Errorf("character is not in ASCII range")
	}

	var lines []string
	// characters in Ascii begin from index 32 so we subtract to get the right indexing
	char := int(C) - 32

	// loop through each line in the character and add it to the lines array

	for i := 0; i <= 7; i++ {
		lines = append(lines, fontData[char][i])
	}

	return lines, nil
}
