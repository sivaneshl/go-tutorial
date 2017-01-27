// methods with pointer receiver type

package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

func (v *vertex) mScale(f float64) { // method that takes a pointer receiver
	v.x = v.x * f
	v.y = v.y * f
}

func fnScale(v *vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v vertex) mAbs() float64 {
	return math.Pow(v.x, v.y)
}

func main() {
	v := vertex{3, 4}
	v.mScale(10) // even though v is a value and not a pointer, the method with the pointer receiver is called automatically
	fmt.Println(v)

	fnScale(&v, 10)
	fmt.Println(v)

	p := &vertex{4, 5} // even though v is a value and not a pointer, the method with the pointer receiver is called automatically
	fmt.Println(p.mAbs())

}
