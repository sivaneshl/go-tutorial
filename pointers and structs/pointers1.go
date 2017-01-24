// pointers1.go

package main 

import (
    "fmt"
    )
    
func main () {
    
    i, j := 10, 20
    
    fmt.Println(i, j)
    
    p := &i // point p to i 
    fmt.Println(*p) // read i through pointer p
    
    *p = 15 // set the value of i through pointer p
    fmt.Println(i)  // new value of i
    
    p = &j  // point p to j
    *p = *p / 5
    fmt.Println(i, j)
    
}    