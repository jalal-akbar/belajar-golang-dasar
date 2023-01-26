package packagess

import (
	"fmt"
	"math"
	"testing"
)

func TestMath(t *testing.T) {
	// Round returns the nearest integer, rounding half away from zero
	fmt.Println("Round 1.7 : ", math.Round(1.7))
	fmt.Println("Round 1.3 : ", math.Round(1.3))
	// Floor returns the greatest integer value less than or equal to x.
	fmt.Println("Floor 1.7 : ", math.Floor(1.7))
	// Ceil returns the least integer value greater than or equal to x.
	fmt.Println("Ceil 1.3 : ", math.Ceil(1.3))
	// Max returns the larger of x or y.
	fmt.Println("Max of 10 and 20 : ", math.Max(10, 20))
	// Max returns the smaller of x or y.
	fmt.Println("Min of 10 and 20 : ", math.Min(10, 20))

}
