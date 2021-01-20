package main

import (
	"fmt"
	"math"
)

func float_equal(x, y, precision float64) bool{
	xi, xf := math.Modf(x)
	yi, yf := math.Modf(y)
	xfi, _ := math.Modf(xf * math.Pow(10, precision))
	yfi, _ := math.Modf(yf * math.Pow(10, precision))

	return xi == yi && xfi == yfi
}

func Sqrt(x, guess, precision float64) (float64, int) {
	z := guess
	last := z
	z -= (z*z - x) / (2 * z)
	i := 0
	for  ;  ! float_equal(last, z, precision) && i < 100 ; i++ {
		//fmt.Println(z)
		last = z
		z -= (z*z - x) / (2 * z)
	}
	return z, i
}

func main() {
	var x float64 = 3
	var guess float64 = x / 2
	precision := 14.

	sqrt, iterations := Sqrt(x, guess, precision)
	fmt.Printf("Sqrt of %v is %v\n", x, sqrt)
	fmt.Println("Number of iterations", iterations)
}
