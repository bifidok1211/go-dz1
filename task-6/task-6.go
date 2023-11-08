package main

import (
	"fmt"
	"strings"
)

func count(arr []string) int {
	u := make(map[string]bool)
	for _, num := range arr {
		u[num] = true
	}
	return len(u)
}

func main() {
	var s string
	fmt.Println("Введите s")
	fmt.Scan(&s)
	w := strings.Join(strings.Split(s, ""), " ")
	m := strings.Fields(w)
	count := count(m)
	fmt.Println(count)

}
