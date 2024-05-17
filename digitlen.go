package piscine

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if n < 1 {
		n = -n
	}
	count := 1
	
	for n > 1 {
		n /= base
		count++
			
	}
	return count
}
