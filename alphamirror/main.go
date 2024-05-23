package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1]

	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	var result string
	for _, char := range args {
		if char >= 'a' && char <= 'z' {
			char = 'z' - char + 'a'
			result += string(char)
		} else if char >= 'A' && char <= 'Z' {
			char = 'Z' - char + 'A'
			result += string(char)
		} else {
			result += string(char)
		}
	}
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}