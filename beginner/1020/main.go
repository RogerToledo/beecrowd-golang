package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)

	anos := num / 365
	meses := (num % 365) / 30
	dias := (num % 365) % 30

	fmt.Printf("%d ano(s)\n", anos)
	fmt.Printf("%d mes(es)\n", meses)
	fmt.Printf("%d dia(s)\n", dias)
}
