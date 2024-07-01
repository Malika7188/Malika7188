package functions

import "fmt"

// function that generates ascii characters using provided rune
func AsciiChar(C rune, fontData [95][8]string) ([]string, error) {
	if C < 32 || C > 126 { // Check if the character is ASCII
		err := fmt.Errorf("character '%c' (%d) is not in ASCII range", C, C)
		return nil, err
	}

	var lines []string // characters in Ascii begin from index 32 so we subtract to get the right indexing
	char := int(C) - 32

	for i := 0; i <= 7; i++ { // loop through each line in the character and add it to the lines array
		lines = append(lines, fontData[char][i])
	}

	return lines, nil
}
