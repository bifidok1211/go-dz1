package main

import (
	"fmt"
)

func main() {
	var a int64
	var b int64
	var c int64
	fmt.Println("Введите а")
	fmt.Scan(&a)
	fmt.Println("Введите b")
	fmt.Scan(&b)
	fmt.Println("Введите c")
	fmt.Scan(&c)
	if (a+b) > c && (b+c) > a && (a+c) > b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
