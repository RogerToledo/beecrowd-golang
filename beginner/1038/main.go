package main

import "fmt"

func main() {
	var x, y int

	fmt.Scanln(&x, &y)

	prod := map[int]float64{
		1: 4.00,
		2: 4.50,
		3: 5.00,
		4: 2.00,
		5: 1.50,
	}

	total := prod[x] * float64(y)

	fmt.Printf("Total: R$ %.2f\n", total)
}