package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	if a := 2; a < 10 {
		a++
		fmt.Println(a)
	}
	mapTest := make(map[string]int)
	mapTest["bob"] = 23
	value, present := mapTest["12"]
	fmt.Println(value, present)
	list := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
	// fizzBuzz()
	deferList()
	// 函数同样可以被赋值给变量
	a := func() {
		fmt.Println("hello")
	}
	a()
	// var xs = map[int]func() int{
	// 	1: func() int { return 10 },
	// 	2: func() int { return 20 },
	// 	3: func() int { return 30 },
	// }

	tt := 3

	var p *int = &tt
	fmt.Printf("tt: %v %x value: %v", p, &tt, tt)
	pp := new(person)
	pp.age = 2
	pp.name = "lzzz"
	fmt.Printf("%v", pp)

}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println(i, "FizzBuzz")
		case i%3 == 0:
			fmt.Println(i, "Buzz")
		case i%5 == 0:
			fmt.Println(i, "Fizz")
		}
	}
}

// defer 的LIFO测试
func deferList() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// 回调函数
func printlit(x int) {
	fmt.Printf("%d", x)
}

func callback(y int, f func(int)) {
	f(y)
}

type person struct {
	name string
	age  int
}
