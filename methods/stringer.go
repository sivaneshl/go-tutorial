//stringer - A Stringer is a type that can describe itself as a string.

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func main() {
	a := Person{"Ronaldo", 33}
	b := Person{"Rooney", 35}
	fmt.Println(a, b)
}
