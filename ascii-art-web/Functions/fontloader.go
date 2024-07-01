package functions

import (
	"bufio"
	"os"
)

// function fontloader to process the fontfile into a 2D array
func Fontloader(fontFile string) ([95][8]string, error) {
	var fontData [95][8]string // initialize fontdata varable to store the ascii-art characters

	file, err := os.Open(fontFile) // opening the font file for reading
	if err != nil {
		return [95][8]string{}, err
	}

	scanner := bufio.NewScanner(file) // create a new scaner to read the file line by line
	CharacterIndex := -1              // initialize a variable to track the current character which begins with a space hence -1
	LineIndex := 0                    // initialize a line index to track the current line being processed

	for scanner.Scan() {
		if scanner.Text() == "" {
			CharacterIndex++
			LineIndex = 0
			continue

		} else { // if the curent charcter not empty it's copied to thr font data variable line by line
			fontData[CharacterIndex][LineIndex] = scanner.Text()
			LineIndex++

		} // if scanner encounters an error it returns an empty array
		if err := scanner.Err(); err != nil {
			return [95][8]string{}, err
		}

	} // if scanner encounters no error it retuns font data with ascii characters in it
	return fontData, nil
}
