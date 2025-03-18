package main

import (
	"fmt"
)

func add(x, y int) int { return x + y }

func swap(x, y string) (string, string) { return y, x }

func main() {
	var _ = add(42, 13)
	var st, ts = swap("this", "is me")
	fmt.Println(st, ts)
	if st == "" {
		panic("something")
	}
}
