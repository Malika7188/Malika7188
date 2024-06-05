package main

import (
	"fmt"
	"os"
	//"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	words := ""
	for _, str := range os.Args[1:] {
		words += str
	}
	result := ""
	seen := make(map[rune]bool)
	for _, char := range words {
		if !seen[char] {
			seen[char] = true
			result += string(char)
		}
	}
	fmt.Println(result)
}
