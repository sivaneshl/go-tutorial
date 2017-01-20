//switch_case.go

package main 

import (
	"fmt"
	"runtime"
)

func main () {
	switch curr_os := runtime.GOOS; curr_os {
		case "darwin":
			fmt.Println("OS X", curr_os)
		case "linux":
			fmt.Println("LINUX", curr_os)	
		default:
			fmt.Println("%s.", curr_os)	
	}
}