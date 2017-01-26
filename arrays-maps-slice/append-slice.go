// append slice

package main

import (
	"fmt"
)

func fnPrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}

func main() {

	var s []int
	fnPrintSlice(s)

	s = append(s, 0)
	fnPrintSlice(s)

	s = append(s, 1)
	fnPrintSlice(s)

	s = append(s, 2, 3, 4, 5) // add more elements at a time
	fnPrintSlice(s)

}
