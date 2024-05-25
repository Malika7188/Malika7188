package goreloaded

import (
	"fmt"
	"strconv"
)

func Hex(words []string) []string {
	for i, result := range words {
		// if there is a string containing (hex) in the array
		if result == "(hex)" || result == "(Hex)" || result == "(HEX)" && i > 0 {
			num, err := strconv.ParseInt(words[i-1], 16, 64)
				if err != nil {
				fmt.Println("Not Hex value")
			}
			words[i-1] = fmt.Sprint(num)
			words = append(words[:i], words[i+1:]...)
		}
	}
	return words
}
