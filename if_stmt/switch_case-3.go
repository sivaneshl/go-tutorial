// switch_case-3.go

package main 

import (
	"fmt"
	"time"
)

func main () {

	cur_time := time.Now()
	fmt.Println (cur_time.Hour(), cur_time.Minute(), cur_time.Second())
	switch {

		case cur_time.Hour() < 12:
			fmt.Println("Good Morning")
		case cur_time.Hour() < 17:
			fmt.Println("Good Afternoon")
		default: 
			fmt.Println("Good Evening")	
	}
}