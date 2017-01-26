// maps word count

package main

import (
	"fmt"
	"strings"
)

func fnWordLen(s string) map[string]int {
	m := make(map[string]int)
	f := strings.Fields(s)
	for i := range f {
		m[f[i]] = len(f[i])
	}
	return m
}

func fnWordCount(s string) map[string]int {
	m := make(map[string]int)
	f := strings.Fields(s)

	for i := range f {
		m[f[i]] = m[f[i]] + 1
	}

	return m
}

func main() {

	s := "I am learning I am Go"
	fmt.Println(fnWordLen(s))
	fmt.Println(fnWordCount(s))

}
