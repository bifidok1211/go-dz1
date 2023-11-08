package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Println("Введите s")
	fmt.Scan(&s)
	replacer := strings.NewReplacer("1", "one")
	out := replacer.Replace(s)
	fmt.Println(out)
}
