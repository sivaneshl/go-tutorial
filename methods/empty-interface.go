// empty interface

package main

import "fmt"

func main() {

	var i interface{}
	fndescribe(i)

	i = 10
	fndescribe(i)

	i = "golang"
	fndescribe(i)

}

func fndescribe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
