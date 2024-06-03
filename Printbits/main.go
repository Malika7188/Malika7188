package main

import (
	"fmt"
	"os"
	//"golang.org/x/text/number"
	// "strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num := Atoi(os.Args[1])

	var bin string

	for num > 0 {
		if num % 2 != 0 {
			bin = "1" + bin
		} else {
			bin = "0" + bin
		}

		num /= 2
	}

	zeros := ""
	if len(bin) < 8 {
		for i := 0; i < 8 - len(bin); i++ {
			zeros += "0"
		}

		bin = zeros+bin
	}

	fmt.Println(bin)
}

func Atoi(str string) int {
	num := 0
	for _, val := range str {
		if val >= '0' && val <= '9' {
			num = (num * 10) + int(val-'0')
		}
	}
	return num
}
