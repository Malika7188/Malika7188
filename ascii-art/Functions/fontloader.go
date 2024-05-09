package functions

import (
	"bufio"
	"os"
)

func Fontloader(fontFile string) ([95][8]string, error) {
	var fontData [95][8]string

	file, err := os.Open(fontFile)
	if err != nil {
		return [95][8]string{}, err
	}

	scanner := bufio.NewScanner(file)
	CharacterIndex := -1
	LineIndex := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			CharacterIndex++
			LineIndex = 0
			continue

		} else {
			fontData[CharacterIndex][LineIndex] = scanner.Text()
			LineIndex++

		}
		if err := scanner.Err(); err != nil {
			return [95][8]string{}, err
		}

	}
	return fontData, nil
}
