package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	fmt.Println("Введите n")
	fmt.Scan(&n)
	a := make([]int, n)
	temp := 0
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(10)
	}
	for i := 0; i < n; i++ {
		fmt.Print(a[i])
	}
	println()
	for i := n; i > 1; i-- {
		temp = a[i-1]
		a[i-1] = a[i-2]
		a[i-2] = temp
	}
	for i := 0; i < n; i++ {
		fmt.Print(a[i])
	}
}
