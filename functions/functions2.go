package main

import (
    "fmt"
    "math"
)

func fnSplit (xy float64) (x, y float64) {
    y = math.Mod (xy, 10) 
    x = xy - y
    return
}

func main () {
    fmt.Println (fnSplit (15))
}
