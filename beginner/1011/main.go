package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = 3.14159
	var r int

	fmt.Scan(&r)
	
	volume := (4.0 / 3) * pi * math.Pow(float64(r), 3)

	fmt.Printf("VOLUME = %.3f\n", volume)
}