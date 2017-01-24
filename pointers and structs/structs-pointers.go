// pointers to structs

package main

import (
	"fmt"
)

type myNewStruct struct {
	x int
	y string
}

func main() {

	v := myNewStruct{1, ""}
	p := &v

	(*p).x = 10 // set value of x using pointer
	p.y = "ten" // it can be simply mentioned as p.y - instead of (*p).y

	fmt.Println(v)

}
