// range of a for

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	for i, v := range pow {
		fmt.Printf("2 power %d = %d \n", i, v)
	}

	for i, v := range pow {
		fmt.Println(pow[i], v)
	}

	// for every iteration range gives 2 values - 1st index, 2nd value

}
