package main

import (
    "fmt"
)

func fnAdd (x int, y int) int {
    return x + y
}

func main () {
    fmt.Println("Sum of 2 numbers: ", fnAdd(3, 4))
}
