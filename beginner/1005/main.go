package main

import "fmt"

func main() {
	const (
		weightA = 3.5
		weightB = 7.5
	)
	var (
		a float32
		b float32
	)

	fmt.Scan(&a)
	fmt.Scan(&b)

	
	updateA := a * weightA
	updateB := b * weightB

	media := (updateA + updateB) / (weightA + weightB)

	fmt.Printf("MEDIA = %.5f\n", media)
}