package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"ascii-art-color/functions"
)

// Map of color names to ANSI escape codes
var colorMap = map[string]string{
	"black":   "30",
	"red":     "31",
	"orange":  "38;5;208",
	"green":   "32",
	"yellow":  "33",
	"blue":    "34",
	"magenta": "35",
	"cyan":    "36",
	"white":   "37",
	"gray":    "90",
	"purple":  "35",
}

func main() {
	useMes := `Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"`

	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println(useMes)
		return
	}

	ar := os.Args[1]
	if strings.HasPrefix(ar, "-") && !strings.HasPrefix(ar, "--") {
		fmt.Println(useMes)
		fmt.Println("EXPECING: --")
		os.Exit(1)
	}
	if strings.HasPrefix(ar, "--") && !strings.Contains(ar, "=") {
		fmt.Println(useMes)
		fmt.Println("EXPECTING: =")
		os.Exit(1)
	}

	// Define and parse the color flag
	color := flag.String("color", "", "Specify color")
	flag.Parse()

	args := flag.Args()
	var substr string
	var inputString string
	fontFile := "standard"

	// Argument handling
	switch len(args) {
	case 1:
		inputString = args[0]
	case 2:
		substr = args[0]
		inputString = args[1]
	case 3:
		substr = args[0]
		inputString = args[1]
		fontFile = args[2]
	default:
		log.Fatal("Invalid number of arguments")
		return
	}

	// Load the font data
	fontArray, err := functions.FontLoader(fontFile)
	if err != nil {
		fmt.Println("Error loading font:", err)
		return
	}

	inputString = strings.ReplaceAll(inputString, "\\n", "\n")
	lines := strings.Split(inputString, "\n")
	var colorCode string
	// Get color code if color flag is specified
	if *color != "" {
		var ok bool
		if strings.HasPrefix(*color, "rgb(") {
			colorCode = functions.RgbToAnsi(*color)
		} else {
			colorCode, ok = colorMap[*color]
			if !ok {
				fmt.Println("Invalid color name")
				return
			}

		}
	}
	count := 0
	for _, line := range lines {
		if line == "" {
			count++
			if count < len(lines) {
				fmt.Println()
				continue
			}
			continue
		}

		concatenatedLines := make([]string, 8)
		var sublines []string

		// Check if substr is specified
		if substr == "" {
			sublines = []string{line}
		} else {
			sublines = strings.Split(line, substr)
		}

		for i, str := range sublines {
			// Process each character in the subline
			for _, char := range str {
				asciiArt, err := functions.PrintChar(char, fontArray)
				if err != nil {
					fmt.Println("Error: Character outside ASCII range encountered")
					return
				}
				for j, artLine := range asciiArt {
					concatenatedLines[j] += artLine
				}
			}

			// Add the colored substring if not the last subline
			if substr != "" && i < len(sublines)-1 {
				coloredSub := functions.ColoredSub(substr, fontArray, colorCode)
				for j, coloredLine := range coloredSub {
					concatenatedLines[j] += coloredLine
				}
			}
		}

		// If substr is not specified, color the entire string
		if substr == "" && colorCode != "" {
			concatenatedLines = functions.ColoredSub(line, fontArray, colorCode)
		}

		// Print the concatenated ASCII art lines
		for _, line := range concatenatedLines {
			fmt.Println(line)
		}
	}
}
