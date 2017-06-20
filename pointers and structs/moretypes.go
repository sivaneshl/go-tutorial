package main

import "fmt"

type Vertex struct {
	x int
	y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{x: 1}
	v3 = Vertex{}
	p2 = &Vertex{10, 20}
)

func main() {
	i, j := 5, 10

	p := &i
	// fmt.Println(p, *p)
	*p = 15
	// fmt.Println(*p, i)

	p = &j
	*p = *p * i
	// fmt.Println(j)

	fmt.Println(i, j, Vertex{i, j})
	v := Vertex{i, j}
	v.x = 100
	fmt.Println(v)

	p1 := &v
	p1.y = 200
	fmt.Println(v, *p1)

	fmt.Println(v1, v2, v3, *p)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes, primes[4])
	var s []int = primes[1:4]
	var s1 []int = primes[0:5]
	fmt.Println(s, s1)

}
