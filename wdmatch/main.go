package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	str1 := args[0]
	str2 := args[1]

	var result string
	// count := 0

	for _, char := range str1 {
		output := false
		for i, char1 := range str2 {
			if char == char1 {
				result += string(char1)
				str2 = str2[i:]

				output = true
				break
			}
		}
		if !output {
			return
		}

	}
	fmt.Println(result)
}
