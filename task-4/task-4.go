package main

import (
	"fmt"
	"math/rand"
)

func bo(a [][]int, n int) bool {
	c := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == a[j][i] {
				c++
			}
		}
	}
	if c == n*n {
		return true
	} else {
		return false
	}
}

func main() {
	var n int
	fmt.Println("Введите n")
	fmt.Scan(&n)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = rand.Intn(2)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(a[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	if bo(a, n) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
