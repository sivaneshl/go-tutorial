// exercise-loops-functions.go

package main

import (
	"fmt"
	"math"
)

func fnNewtonSqrt(x float64) float64 {
	// initialize z as float64
	z := 1.0 // z = z(n)
	y := 1.0 // y = z(n+1)

	if z == 1.0 {
		y = z - (math.Pow(z, 2)-x)/(2*z)
	}
	z = y + 1 // hack to make z > y

	for z > y {
		z = y
		y = z - (math.Pow(z, 2)-x)/(2*z)
	}

	/*
		for i < 10 {
			y = z - (math.Pow(z, 2) - x) / (2 * z)
			fmt.Println(z, y)
			z = y
			i += 1
		}
	*/
	return y

}

func main() {
	fmt.Println(fnNewtonSqrt(20))
}
