package main

import "fmt"

func fnAdd (x, y int) (string, int) {
    return ("X + Y = "), (x + y) 
    
}

func main (){
    a, b := fnAdd(4, 5)
    fmt.Println(a, b)
    
}
