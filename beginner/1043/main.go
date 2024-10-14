package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Scan(&a, &b, &c)

	ab := math.Abs(a - b)
	bc := math.Abs(b - c)
	ca := math.Abs(c - a)

	if ab < c && c < a + b {
		if bc < a && a < b + c {
			if ca < b && b < a + c {
				fmt.Printf("Perimetro = %.1f\n", a + b + c)
			}
		}
	} else {
		fmt.Printf("Area = %.1f\n", ((a + b) / 2) * c)
	}
}