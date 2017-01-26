// range

package main

import "fmt"

func main() {

	pow := make([]int, 10)

	for i := range pow { // omit value of range and store index only
		pow[i] = i ^ 2
	}

	for i, v := range pow { // omit index and store value only
		fmt.Println(i, v)
	}

}
