// slices

package main

import "fmt"

func main() {
	cVowels := [5]string{"a", "e", "i", "o", "u"}

	var s []string = cVowels[1:4] // slice is a subset of an array.. decalre first..  type []T is a slice with elements of type T

	s1 := cVowels[2:4] // s1[0] will be s[2]
	fmt.Println(s, s1)

	s1[0] = "XXX"      // modify the elements of the slice
	fmt.Println(s, s1) // will update the underlying array

	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v \n", len(s1), cap(s1), s1)

	var s_nil []int
	fmt.Println(s_nil, len(s_nil), cap(s_nil))

}
