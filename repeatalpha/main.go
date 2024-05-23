package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1]

	// if len(args) != 2 {
	// 	return
	// }
	//count := 1
	// var result string
	count := 1
	for _, char := range args {
		if char >= 'a' && char <= 'z' {
			count = int(char - 'a' + 1)
			// result += string(char)
		} else if char >= 'A' && char <= 'Z' {
			count = int(char - 'A' + 1)
			// result += string(char)
		}
		for i := 0; i < count; i++ {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
