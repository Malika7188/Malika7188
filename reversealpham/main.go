package main

import "github.com/01-edu/z01"

func main() {
	// str := "zyxwvutsrqponmlkjihgfedcba"

	// for _, char := range str {
	// 	if char%2 == 0 {
	// 		z01.PrintRune(char - 32)
	// 		continue
	// 	}
	// 	z01.PrintRune(char)
	// }
	// z01.PrintRune('\n')

	for i := 'a'; i <= 'z'; i++ {
		if i%2 == 0 {
			z01.PrintRune(i - 32)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
