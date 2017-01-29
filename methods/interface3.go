// Under the covers, interface values can be thought of as a tuple of a value and a concrete type:

// (value, type)

// An interface value holds a value of a specific underlying concrete type.

// Calling a method on an interface value executes the method of the same name on its underlying type.

package main

import (
	"fmt"
)

type myfloat float64

func (fvar myfloat) mymethod() {
	fmt.Println(fvar)
}

type mystruct struct {
	s string
}

func (svar *mystruct) mymethod() {
	if svar == nil {
		fmt.Println("NIL")
		return
	}
	fmt.Println(svar.s)
}

type myinterface interface {
	mymethod()
}

func main() {

	var ivar myinterface

	ivar = myfloat(1.432)
	fmt.Printf("(%v, %T)\n", ivar, ivar)
	ivar.mymethod()

	ivar = &mystruct{"glang"}
	fmt.Printf("(%v, %T)\n", ivar, ivar)
	ivar.mymethod()

	var svar *mystruct
	ivar = svar
	fmt.Printf("(%v, %T)\n", ivar, ivar)
	ivar.mymethod()

}
