package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	fmt.Println("Введите а")
	fmt.Scan(&a)
	fmt.Println("Введите b")
	fmt.Scan(&b)
	c := math.Pow(a*a+b*b, 0.5)
	fmt.Println("Гипотенуза")
	fmt.Println(c)

}
