package main

import "fmt"

func SwapBits(octet byte) byte {
	s1 := octet >> 4
	s2 := octet << 4

	return s1 | s2
}

func main() {
	fmt.Println(SwapBits(7))
}
