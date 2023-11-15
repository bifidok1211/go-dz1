package main

import (
	"fmt"
	"strings"
)

type tank struct {
}

func count(arr []string) int {
	u := make(map[string]struct{})
	for _, num := range arr {
		if _, ok := u[num]; !ok {
			u[num] = struct{}{}
		}
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
