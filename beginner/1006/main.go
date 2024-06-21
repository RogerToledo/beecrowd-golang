package main

import "fmt"

func main() {
	const (
		weightA = 2
		weightB = 3
		weightC = 5
	)
	var a, b, c float32

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	
	updateA := a * weightA
	updateB := b * weightB
	updateC := c * weightC

	media := (updateA + updateB + updateC) / (weightA + weightB + weightC)

	fmt.Printf("MEDIA = %.1f\n", media)
}