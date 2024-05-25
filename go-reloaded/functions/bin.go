package goreloaded

import (
	"fmt"
	"strconv"
)

func Bin(words []string) []string {
	for i, result := range words {
		// if there is a string containing (hex) in the array
		if result == "(bin)" || result == "(Bin)" || result == "(BIN)" && i > 0 {
			num, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				fmt.Println("Not bin value")
			}
			words[i-1] = fmt.Sprint(num)
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
