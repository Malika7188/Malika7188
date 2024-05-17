package mathskils

import (
	"math"
	"testing"
)

func TestHandleMedian(t *testing.T) {
	input := []float64{1, 2, 3, 4, 5, 6}
	expected := float64(4)
	got := HandleMedian(input)

	if math.Round(got) != expected {
		t.Fatalf("expected %v and got %v", expected, math.Round(got))
	}
}
