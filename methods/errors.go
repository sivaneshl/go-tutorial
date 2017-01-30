// errors

package main

import (
	"fmt"
	"strconv"
)

func main() {

	i, err := strconv.Atoi("422d")
	if err != nil {
		fmt.Printf("ERROR: couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)

}
