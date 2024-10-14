package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	//fmt.Scanln(&a, &b, &c)

	a = 10.0
	b = 20.1
	c = 5.1

	delta := math.Pow(b,2) - 4 * a * c
	div := 2 * a

	if delta < 0 || div == 0 {
		r1 := (-b + math.Sqrt(delta)) / div
		r2 := (-b - math.Sqrt(delta)) / div

		fmt.Printf("R1 = %.5f\n", r1)
		fmt.Printf("R2 = %.5f\n", r2)
	} else {
		fmt.Println("Impossivel calcular")
	}
}