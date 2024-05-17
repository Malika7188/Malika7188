package mathskils

import (
	"math"
	"testing"
)

func TestVariance(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := float64(7)
	got := HandleVariance(input)

	if math.Round(got) != expected {
		t.Fatalf("expected %v and got %v", expected, math.Round(got))
	}
}
