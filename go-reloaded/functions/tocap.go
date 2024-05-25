package goreloaded

import (
	"fmt"
	"strconv"
	"strings"
)

// converting to upper
func Cap(words []string) []string {
	for i, result := range words {
		if strings.Contains(strings.ToLower(result), "(cap,") {
			num, err := strconv.Atoi(words[i+1][:len(words[i+1])-1]) // read the argument content to check the errors then panic if theres any error found
			if err != nil {
				panic(err)
			}
			if num > i { // checks if the number is greater than the word at words[i], then loops from thge first index 0
				for j := 0; j <= i; j++ {
					words[j] = strings.Title(words[j])
				}
				words[i] = " "
				words[i+1] = " "
			} else if num < 0 { // check if theres an negative number then prints the error
				fmt.Println("The input in invalid")
			} else if num <= i {
				for j := i - num; j < i; j++ {
					words[j] = strings.Title(words[j])
				}
				words[i] = " "
				words[i+1] = " "
			}
		} else if strings.Contains(strings.ToLower(result), "(cap)") { // check if the word result is (up) an is greater then 0
			words[i-1] = strings.Title(words[i-1])
			words[i] = " "
		}
	}
	words = RemoveSpaces(words)
	return words
}
