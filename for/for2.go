package main

import "fmt"

func main () {
	
	i := 0

	// for ; i < 20 ; --> dropping semi-colons makes it similar to while
	for i < 20 {
		i = i + 2	
		fmt.Println(i)
	}

}

