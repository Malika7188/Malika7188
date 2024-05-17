package mathskils

import (
	"math"
	"testing"
)

func TestStandardDev(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5, 6, 7, 8,9}
	expected := float64(3)
	got := HandleStandardDeviation(input)

	if math.Round(got) != expected {
		t.Fatalf("expected %v and got %v", expected, math.Round(got))
	}
}
