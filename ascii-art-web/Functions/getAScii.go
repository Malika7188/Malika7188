package functions

import (
	//"fmt"
	"log"
	"strings"
)

func Ascii(input string, fontdata string) ([]string, error) {
	file, err := Fontloader(fontdata)
	if err != nil {
		log.Fatal(err)
	}

	input = strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(input, "\n")

	var asciilines []string
	count := 0
	for _, line := range lines {
		concatLines := make([]string, 8)
		if line == "" {
			count++
			if count < len(lines) {
				concatLines = append(concatLines, "")
				continue
			}
			continue
		}
		for _, char := range line {
			asciiArt, err := AsciiChar(char, file)
			if err != nil {
				log.Fatal(err)
			}
			for i, artline := range asciiArt {
				concatLines[i] += artline
			}
		}
		for _, art := range concatLines {
			asciilines = append(asciilines, art)
		}

	}
	return asciilines, nil
}
