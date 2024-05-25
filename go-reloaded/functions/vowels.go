package goreloaded

// checks is s string
func CheckVowels(s byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H'} //initializing the vowels
	for _, val := range vowels { //range through the vowels
		if val == s {
			return true
		}
	}
	return false
}
