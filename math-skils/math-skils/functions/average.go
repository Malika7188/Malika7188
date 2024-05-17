package mathskils

func HandleAverage(number []float64) float64 {
	var sum float64
	Total := len(number)

	if number == nil {
		return 0.0
	}
	for _, value := range number {
		sum += value
	}
	return sum / float64(Total)
}
