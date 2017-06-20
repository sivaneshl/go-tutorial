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
	s[2] = 100
	fmt.Println(primes, s, s1)

	s1 = s1[:0]
	fmt.Println(len(s1), cap(s1), s1)
	s1 = s1[:4]
	fmt.Println(len(s1), cap(s1), s1)

	s2 := make([]int, 0, 5)
	fmt.Println(len(s2), cap(s2), s2)
	s2 = append(s2, 0, 1, 2)
	fmt.Println(len(s2), cap(s2), s2)

	for i, j = range primes {
		fmt.Println(i, j)
	}

	pow := make([]int, 10)
	for i = range pow {
		pow[i] = i * i
	}
	for _, j = range pow {
		// fmt.Println(j)
	}

	var m map[int]string
	m = make(map[int]string)
	m[1] = "one"
	fmt.Println(m[1])

	var m1 map[int]Vertex
	m1 = make(map[int]Vertex)
	m1[1] = Vertex{1, 11}
	m1[2] = Vertex{2, 21}
	fmt.Println(m1)

	var m2 = map[int]Vertex{
		1: Vertex{1, 11},
		2: Vertex{2, 21},
		3: Vertex{3, 31},
	}

	m2[4] = Vertex{4, 41}
	elem, ok := m2[4]

	fmt.Println(m2, elem, ok)

	delete(m2, 4)
	elem, ok = m2[4]
	fmt.Println(m2, elem, ok)

}
