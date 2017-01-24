// structs

package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

type myNewStruct struct {
	i int
	x string
	l bool
}

func main() {

	v := Vertex{1, 2}
	fmt.Println(v.x, v.y) // Struct fields are accessed using a dot.

	fmt.Println(Vertex{1, 2})

	fmt.Println(myNewStruct{5, "five", false})

}
