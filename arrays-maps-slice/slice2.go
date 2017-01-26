package main

import "fmt"

func fnMake2D (x,y int) [][]int {
    s1 := make([][]int, x)

	for i := 0; i < x; i++ {
		s1[i] = make([]int, y)
		for j := 0; j < y; j++ {
			s1[i][j] = (i ^ j)
		}
	}
    return s1
}

func main() {

	fmt.Println(fnMake2D(2,4))

}
