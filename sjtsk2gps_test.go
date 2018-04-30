package sjtsk2gps

import (
	"testing"
	"math"
	"fmt"
)

func TestConvert(t *testing.T) {
	la, lo, _ := Convert(-515244.88, -1166620.04 , 0)

	fmt.Printf("%f x %f\n", la, lo)

	equalWithPrecision(t, la, float64(49.21717974877753))
	equalWithPrecision(t, lo, float64(17.74835130863854))
}

func equalWithPrecision(t *testing.T, a float64, b float64) {
	if math.Abs(a - b) > 0.000001 {
		t.Fatal(fmt.Printf("%f != %f", a, b))
	}
}

