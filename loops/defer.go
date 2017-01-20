//defer

package main

import "fmt"

func main() {
	fmt.Println("counting")

	// defer function will execute after the surrounding function executes
	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order. 

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
