package main

import (
	"fmt"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	fmt.Println("Hello World")
	i, j := 42, 2071
	p := &i
	fmt.Println(p, *p)
	*p = 21
	fmt.Println(i)
	fmt.Println(j)
	v := Vertex{1, 2}
	fmt.Println(v.x, v.y)
	v.x = 33
	fmt.Println(v.x, v.y)
	pp := &v
	pp.x = 1e9
	fmt.Println(v.x, v.y)
	lit := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(lit[:3])
	fmt.Println(lit[1:4])
	fmt.Println(lit[2:])
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Map
	mothdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
	}
	fmt.Println(mothdays)
}

type Vertex struct {
	x int
	y int
}
