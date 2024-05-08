package mathskils

import (
	"sort"
)

func HandleMedian(number []float64) float64 {
	sort.Float64s(number) // sorting the numbers
	midnumber := len(number)

	if number == nil {
		return 0.0 // checks if the package containing numbers is empty
	}
	if midnumber%2 == 0 { // checks if the number is even and return the average of the 2 middle numbers
		return (number[midnumber/2-1] + number[midnumber/2]) / 2
	} else {
		return number[midnumber/2] // if the number is odd, it returs the middle numbe
	}
}
