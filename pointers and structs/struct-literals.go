// struct literals

package main

import "fmt"

type Vertex struct {
	x int
	y string
}

func main() {

	v1 := Vertex{1, "one"}
	fmt.Println(v1)

	v2 := Vertex{x: 1} // assign value of v2.x. v2.y takes default value
	fmt.Println(v2)

	v3 := Vertex{} // both x and y gets default values
	fmt.Println(v3)

	p := &Vertex{10, "ten"} // set pointer to struct
	fmt.Println(p, *p)      // The special prefix & returns a pointer to the struct value.
}
