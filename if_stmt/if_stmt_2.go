//if_stmt_2.go

package main

import (
	"fmt"
	"math"
)

var (
	lim float64 = 10
)

func fnPower (x,y float64) float64 {
	if z := math.Pow(x, y); z < lim {
		return z
	} else {
		return math.Mod(z, lim) // variabeles defined in IF stmt can be used within the If...ELSE stmt
	}
}

func main () {

	fmt.Println (fnPower (3,4))

}