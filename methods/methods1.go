// methods

package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

type myfloat float64

func (v vertex) mAbs() float64 { // method mAbs with receiver arg v of type vertex
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func fnAbs(v vertex) float64 { // same method as a function
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (f myfloat) mSqr() myfloat { // method with non-struct type
	return (f * f)
}

func main() {

	v := vertex{3, 4}
	fmt.Println(v.mAbs())
	fmt.Println(fnAbs(v))

	f := myfloat(5)
	fmt.Println(f.mSqr())

}
