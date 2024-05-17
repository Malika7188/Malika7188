package mathskils

func HandleVariance(number []float64) float64 {
	var sum float64
	avg := HandleAverage(number) // call average to to use its square diffrence
	p := len(number)

	if p == 0 { // checks if the string is empty then return 0
		return 0.0
	}
	for _, char := range number { // range over the sring then calculate the square difference of average
		sum += (char - avg) * (char - avg)
	}
	return sum / float64(p)
}
