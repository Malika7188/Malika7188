package main

import (
	"fmt"
	"os"
	"strings"

	"fs/functions"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}

	args := os.Args[1:]

	inputString := args[0]
	var fontFile string

	if len(args) == 2 {
		fontFile = args[1]
	} else {
		fontFile = "standard"
	}

	// input := flag.Args()
	// inputString := strings.Join(input, " ")

	// Load the font data
	fontArray, err := functions.Fontloader(fontFile)
	if err != nil {
		fmt.Println("Error loading font:", err)
		return
	}
	inputString = strings.ReplaceAll(inputString, "\\n", "\n")

	//  Print the horizontally concatenated ASCII art representation of each word in the input string
	words := strings.Split(inputString, "\n")
	count := 0
	for _, word := range words {
		if word == "" {
			count++
			if count < len(words) {
				fmt.Println()
				continue
			} else {
				continue
			}
		}
		// initialize an empty array to hold the lines of concatenated Ascii art for the word
		concatenatedLines := make([]string, 8)
		// Concatenate the Ascii art representation of each character in the word horizontally
		for _, char := range word {
			// retriev the Ascii art representation for the character using the AsciiChar function
			asciiArt, err := functions.AsciiChar(char, fontArray)
			if err != nil {
				fmt.Println("Error:Character outside ascii range encountered")
				return
			}

			for i, line := range asciiArt {
				concatenatedLines[i] += line
			}
		}
		// Print the concatenated ASCII art lines for the word
		for _, line := range concatenatedLines {
			fmt.Println(line)
		}

	}
}
