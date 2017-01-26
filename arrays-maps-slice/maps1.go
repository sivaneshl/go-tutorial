// maps

package main

import "fmt"

type coordinates struct {
	fLat, fLong float64
}

func main() {

	var m map[int]string // key int, value string

	m = make(map[int]string) // creating the map - The make function returns a map of the given type, initialized and ready for use.
	m[1] = "one"

	fmt.Println(m[1])

	/*	m1 := make(map[string]coordinates)
		m1["chennai"] = coordinates{13.067439, 80.237617} */

	m1 := map[string]coordinates{
		"chennai":  coordinates{13.067439, 80.237617},
		"banglore": coordinates{12.972442, 77.580643},
		"madurai":  {9.939093, 78.121719}, // you can also omit the type name
	}

	fmt.Println(m1["chennai"], m1["madurai"])
}
