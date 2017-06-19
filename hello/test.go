package main

import (
	"fmt"
	"math"
)

func myfunc(x float64, y float64) float64 {
	return math.Sqrt(x) * (y / math.Pi)
}

func add(x, y float64) float64 {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

/*
func split(z int) (x, y int) {
	x = z * 4 / 9
	y = z - x
	return x, y
}
*/

func split(z int) (int, int) {
	var x, y int
	x = z * 4 / 9
	y = z - x
	return x, y
}

func convert() {
	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
	fmt.Printf("%T\n", f)
}

func main() {
	fmt.Println(myfunc(3, 5))
	fmt.Println(add(myfunc(3, 5), 2*math.Pi))

	a, b := (swap("Hi", "Bye"))
	fmt.Println(a, b)

	fmt.Println(split(17))

	convert()
}
