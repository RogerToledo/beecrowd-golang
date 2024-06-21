package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = 3.14159
	var r float64

	fmt.Scan(&r)

	area := pi * (math.Pow(r, 2))

	fmt.Printf("A=%.4f\n", area)
}