package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//args := os.Args
	if len(os.Args) != 4 {
		fmt.Println("usage: go run . inputst")
	}
	// val1 := args[1]
	// val2 := args[3]

	val1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("error passing file", os.Args[1])
		return
	}
	val2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		return
	}
	operator := os.Args[2]
	operators := []string{"+", "-", "/", "*", "%"}
	op := false

	for _, value := range operators {
		if operator == value {
			op = true
			break			
		} else {
			op = false
		}

	}
	if !op {
		return
	}
	var answer int
	if operator == "+" {
		answer = val1 + val2
	}
	if operator == "-" {
		answer = val1 - val2
	}
	if operator == "*" {
		answer = val1 * val2
	}
	if operator == "/" {
		answer = val1 / val2
	}
	if operator == "%" {
		answer = val1%val2
	}
	if operator == "%" && val2 == 0 {
		os.Stdout.WriteString("No modulo by 0")
		return
	}
	if operator == "/" && val2 == 0 {
		os.Stdout.WriteString("No division by 0")
		return
	}
	if val1 > 9223372036854775807 || val1 < -9223372036854775807 {
		return
	}
	if val2 > 9223372036854775807 || val2 < -9223372036854775807 {
		return
	}
	num := strconv.Itoa(answer)
	os.Stdout.WriteString(num + "\n")

}