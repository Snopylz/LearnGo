package main

import (
	"fmt"
	"math"
	"math/rand"
)

// var c, python, java bool

var i, j int = 1, 2

func main() {
	// var i int
	var python, java = true, "no!"
	fmt.Println("My favorite number is ", rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println(add(2, 3))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(222))
	fmt.Println(i, j, python, java)
	// fmt.Println(i, c, python, java)
}

// func add(a int, b int) int {
// 	return a + b
// }
// x, y int 可以缩写为 x,y int
func add(a, b int) int {
	return a + b
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
