// type switch

package main

import "fmt"

func fnTypeCheck(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Type int <%T>\n", v)
	case string:
		fmt.Printf("Type string <%T>\n", v)
	default:
		fmt.Printf("Type default <%T>\n", v)
	}
}

func main() {

	fnTypeCheck(10)
	fnTypeCheck("golang")
	fnTypeCheck(true)

}
