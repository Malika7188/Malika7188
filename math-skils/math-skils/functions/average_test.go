package mathskils

import (
	"math"
	"testing"
)

func TestAverage(t *testing.T) {
	input := []float64{1,2,3,4,5,6,7,8,9}
	expected := float64(5)
	output := HandleAverage(input)

	if math.Round(output) != expected {
		t.Fatalf("expected %v and output %v", expected, math.Round(output))
	}
}