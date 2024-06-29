package main

import (
	//"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num := Atoi(os.Args[1])
	result := ""
	for i := 1; i <= 9; i++ {
		ans := i * num
		res := Itoa(ans)
		result += Itoa(i) + " * " + Itoa(num) + " = " + res + "\n"
	}
	PrintStr(result)
}
func PrintStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}
func Atoi(str string) int {
	var isneg bool
	if str[0] == '-' {
		str = str[1:]
		isneg = true
	}
	if str[0] == '+' {
		str = str[1:]
	}
	var num int
	for _, val := range str {
		if val >= '0' && val <= '9' {
			num = (num * 10) + int(val-'0')
		} else {
			return 0
		}
	}
	if isneg {
		num *= -1
	}
	return num
}
func Itoa(num int) string {
	isneg := false
	if num < 0 {
		isneg = true
		num *= -1
	}
	result := ""
	for num > 0 {
		digit := num%10
		result = string(digit+'0') + result
		num /= 10

	}
	if isneg {
		return "-" + result
	}
	return result
}
