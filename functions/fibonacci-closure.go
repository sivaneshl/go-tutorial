package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := 1
	j := 0
	sum := i + j
	return func() int {
		sum = i + j
		i = j
		j = sum
		//fmt.Println(i, j, sum)
		return i
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
