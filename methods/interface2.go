// interfaces are implemented implicitly

package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {

	var i I = T{"hello"} //  A value of interface type can hold any value that implements those methods. - here interface I holds the value of type T which implements the method M in I
	i.M()
}
