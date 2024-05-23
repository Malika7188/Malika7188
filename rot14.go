package revision

func Rot14(s string) string {
	result := []rune{}
	var newchar rune

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			newchar = (char-'a'+14)%26 + 'a'
			result = append(result, newchar)
		} else if char >= 'A' && char <= 'Z' {
			newchar = (char-'A'+14)%26 + 'A'
			result = append(result, newchar)
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}
