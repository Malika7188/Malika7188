package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 || len(args) == 0 {
		return
	}
	args1 := args[0]

	var result string

	for i := len(args1) - 1; i > 0; i-- {
		if result == "" && string(args1[i]) == " " {
			continue
		}
		if result != "" && string(args1[i]) == " " {
			break
		}
		result = string(args1[i]) + result
	}
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
