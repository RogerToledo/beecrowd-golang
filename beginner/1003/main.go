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

	soma := a + b

	fmt.Printf("SOMA = %d\n", soma)
}