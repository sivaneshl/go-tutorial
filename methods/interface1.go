// interfaces - An interface type is defined as a set of method signatures.

package main

import (
	"fmt"
	"math"
)

type myfloat float64

type vertex struct{ x, y float64 }

func (v *vertex) mAbs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (f myfloat) mAbs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

type abser interface {
	mAbs() float64
}

func main() {

	var a abser
	f := myfloat(-math.Sqrt2)
	v := vertex{3, 4}

	a = f //  A value of interface type can hold any value that implements those methods.
	fmt.Println(a.mAbs())

	a = &v //  A value of interface type can hold any value that implements those methods.
	fmt.Println(a.mAbs())

}
