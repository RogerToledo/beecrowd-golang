package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64

	fmt.Scanln(&x1, &y1)
	fmt.Scanln(&x2, &y2)

	x := math.Pow(x2 - x1, 2)
	y := math.Pow(y2 - y1, 2)

	distance := math.Sqrt(x + y)

	fmt.Printf("%.4f\n", distance)
}