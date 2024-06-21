package main

import "fmt"

func main() {
	var a, b, c, d int64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)

	diff := (a * b) - (c * d)

	fmt.Printf("DIFERNCA = %d\n", diff)
}