// arrays

package main

import "fmt"

func main() {

	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	fmt.Println(a[0], a[1])

	cVowels := [5]string{"a", "e", "i", "o", "u"} // array literal
	fmt.Println(cVowels)

}
