package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	args1 := args[0]

	word := ""
	var words []string

	for i, letter := range args1 {
		if letter == ' ' || i == len(args1)-1 {
			if i == len(args1)-1{
				word += string(letter)
			}
			words = append(words, word)
			word = ""
		} else {
			word += string(letter)
		}
	}
	fmt.Println(words[0])
}

// if len(args) != 1 || len(args) == 0 {
// 	return
// }
// var result string

// for i := 0; i < len(args); i++ {
// 	if result != "" && string(args1) == " " {
// 		break
// 	}
// 	result = string(args1[i]) + result
// }
// fmt.Println(result)
