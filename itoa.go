package revision

// import "fmt"

// import "golang.org/x/text/number"
func Itoa(number int) string {
	str := ""

	for number > 0 {
		digit := number % 10
		strdigit := string('0' + rune(digit))
		str = strdigit + str
		number /= 10
	}
	return str
}
