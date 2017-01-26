// functions as values

package main

import (
	"fmt"
	"math"
)

func fnCompute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {

	fnHypot := func(x, y float64) float64 {
		return math.Sqrt(x*y + x*y)
	}
	//fmt.Println(fnHypot(3, 4))

	fmt.Println(fnCompute(fnHypot))
	fmt.Println(fnCompute(math.Pow))

}
