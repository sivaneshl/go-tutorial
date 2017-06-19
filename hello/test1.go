package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	if x > 100 {
		return math.Sqrt(x)
	}
	return math.Pi * x
}

func pow(x, y, z float64) float64 {
	if v := math.Pow(x, y); v <= z {
		return v
	} else {
		return v - z
	}
}

func newsqrt(x float64) float64 {
	var z, z1, y float64 = 1, 1, 1

	for i := 1; i <= 10; i++ {
		z1 = z - (math.Pow(z, 2)-x)/(2*z)
		// 		fmt.Println(z, z1)
		if y == z1 {
			return z1
		} else {
			y = z
			z = z1
		}
	}
	return z1
}

func switchcase(x float64) string {
	switch mod5 := math.Mod(x, 5); mod5 {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	default:
		return "ZERO"
	}
}

func deferfunc() {
	fmt.Println("Start")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("End")
}

func main() {
	var i int
	sum := 0
	for i = 1; i <= 10; i++ {
		sum += i
	}

	fmt.Println(sum, sqrt(float64(sum)))

	for sum < 1000 {
		sum += sum
	}

	defer fmt.Println("Hello world!!")

	fmt.Println(sum, sqrt(float64(sum)))

	fmt.Println(pow(3, 2, 10), pow(3, 3, 10))

	fmt.Println(newsqrt(2))

	fmt.Println(switchcase(8))

	deferfunc()

}
