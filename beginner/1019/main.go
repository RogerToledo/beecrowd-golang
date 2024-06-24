package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Scan(&num)

	h := num / 3600
	m := (num % 3600) / 60
	s := (num % 3600) % 60

	fmt.Printf("%d:%d:%d\n", h, m, s)
}