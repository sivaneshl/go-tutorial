package main

import "fmt"

func fnConv (x int) (float64) {
	return float64(x)
}

func main () {
	fmt.Println (fnConv (123))
}
