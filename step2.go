package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	sum := 1
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	for sum < 1000 {
		sum += sum
	}
	// for {

	// }死循环

	// 由这个语句定义的变量的作用域仅在 if 范围之内。
	fmt.Println("SUM:", sum)
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(2, 2, 10),
		pow(3, 3, 20),
	)
	println("Sqrt: ", sqrt(2.0), math.Sqrt(2))
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Print("Linux")
	default:
		fmt.Printf("%s.", os)
	}
	saytime()
}

// func sqrt(x float64) string {
// 	if x < 0 {
// 		return sqrt(-x) + "i"
// 	}
// 	return fmt.Sprint(math.Sqrt(x))
// }

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func sqrt(x float64) float64 {

	z := float64(1)
	for i := 0; i < 10; i++ {
		z = (z - (z*z-x)/(2*z))
	}
	return z
}

func saytime() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Moring！")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
