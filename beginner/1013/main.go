package main

import "fmt"

func main() {
	var a, b, c, maior int

	fmt.Scanln(&a, &b, &c)

	if a > b {
		maior = a
	} else {
		maior = b
	}

	if c > maior {
		maior = c
	}

	fmt.Printf("%d eh o maior\n", maior)
}