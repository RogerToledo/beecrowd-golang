package main

import (
	"fmt"
)

func main() {
	var (
		a int64
		b int64
	)	

	fmt.Scan(&a)
	fmt.Scan(&b)

	prod := a * b

	fmt.Printf("PROD = %d\n", prod)
}