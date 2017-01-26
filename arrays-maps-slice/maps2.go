// maps

package main

import "fmt"

func main() {

	m := make(map[int]string)

	m[1] = "one"
	fmt.Println(m[1])

	m[1] = "two"
	fmt.Println(m[1])

	delete(m, 1)
	fmt.Println(m[1])

	val, ok := m[1]
	fmt.Println(val, ok)

}
