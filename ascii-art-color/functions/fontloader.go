package functions

import (
	"bufio"
	"os"
)
// FontLoader loads font data from a specified text file. The function expects the file name without an extension
func FontLoader(s string) ([95][8]string, error) {
	// Append ".txt" to the provided string to form the filename
	s += ".txt"

	// Initialize a 2D array to hold the font data
	var fontData [95][8]string

	// Open the file for reading
	file, err := os.Open(s)
	if err != nil {
		// Return an empty fontData array and the error if the file cannot be opened
		return [95][8]string{}, err
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	currentCharacter := -1 // Index for the current character being processed
	currentLine := 0       // Index for the current line of the current character

	// Loop through each line in the file
	for scanner.Scan() {
		// If an empty line is encountered, it indicates the start of a new character
		if scanner.Text() == "" {
			currentCharacter++ // Move to the next character
			currentLine = 0    // Reset the line index for the new character
			continue
		} else {
			// Store the line of font data in the appropriate location in the array
			fontData[currentCharacter][currentLine] = scanner.Text()
			currentLine++ // Move to the next line for the current character
		}

		// Check for any errors that occurred during scanning
		if scanner.Err() != nil {
			// Return an empty fontData array and the error if scanning failed
			return [95][8]string{}, scanner.Err()
		}
	}

	// Return the populated fontData array and a nil error to indicate success
	return fontData, nil
}
