package main

import "fmt"

func main() {
	var x, y float64

	fmt.Scanln(&x, &y)

	switch {
	case x > 0 && y > 0:
		fmt.Println("Q1")
	case x < 0 && y > 0:
		fmt.Println("Q2")
	case x < 0 && y < 0:
		fmt.Println("Q3")
	case x > 0 && y < 0:
		fmt.Println("Q4")
	case x == 0 && y != 0:
		fmt.Println("Eixo Y")
	case x != 0 && y == 0:
		fmt.Println("Eixo X")
	default:
		fmt.Println("Origem")	
	}
}
