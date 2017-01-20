//switch_case-2.go

package main 

import (
	"fmt"
	"time"
)

func main () {

	dtcurrday := time.Now().Weekday()

	switch time.Monday {
		case dtcurrday + 0:
			fmt.Println ("TODAY")
		case dtcurrday + 1:
			fmt.Println ("TOMORROW")	
		case dtcurrday + 2:
			fmt.Println ("In two days")
		default:
			fmt.Println ("Too far away")		

	}

	fmt.Println (dtcurrday, (dtcurrday - 1), time.Monday)

}