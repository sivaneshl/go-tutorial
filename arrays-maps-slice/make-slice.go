// make -slice

package main

import "fmt"

func main() {

	a := make([]int, 5) // len(a) = 5
	fnPrintSlice("a", a)

	b := make([]int, 0, 5) //len(b)=5 cap(b)=5
	fnPrintSlice("b", b)

	c := b[:2]
	fnPrintSlice("c", c)

	d := b[2:cap(b)]
	fnPrintSlice("d", d)

}

func fnPrintSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v \n", s, len(x), cap(x), x)
}
