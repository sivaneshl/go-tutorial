package main

import (
	"fmt"
	"math"
)

func fnOddEven (ipiNum int) string {
	if math.Mod(float64(ipiNum), float64(2)) == 0  {
		return "EVEN"
	} else {
		return "ODD"	
	} 

//return int(math.Mod(float64(ipiNum), float64(2)))
}

func main () {

	var iNum int
	iNum = 13
	
	fmt.Println("The ", iNum, " Is ", fnOddEven(iNum))	
}
