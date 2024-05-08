package mathskils

import "math"

func HandleStandardDeviation(number []float64) float64 {
	p := HandleVariance(number)//call variance to this function

	if len(number) == 0 { //checks if the string is empty
		return 0.0
	}
	return math.Sqrt(p) //to calculate standard deviation:get the squareroot of variance
}
