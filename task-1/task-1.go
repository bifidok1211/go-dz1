package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Println("Введите а")
	fmt.Scan(&a)
	fmt.Println("Введите b")
	fmt.Scan(&b)
	c := math.Hypot(a, b)
	fmt.Println("Гипотенуза")
	fmt.Println(c)

}
