package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Println("Введите s")
	fmt.Scan(&s)
	out := strings.ReplaceAll(s, "1", "one")
	fmt.Println(out)
}
