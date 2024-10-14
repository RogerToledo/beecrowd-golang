package main

import "fmt"

func main() {
	var num float64

	fmt.Scan(&num)

	switch {
	case num >= 0 && num <= 25:
		fmt.Println("Intervalo [0,25]")
	case num > 25 && num <= 50:
		fmt.Println("Intervalo (25,50]")
	case num > 50 && num <= 75:
		fmt.Println("Intervalo (50,75]")
	case num > 75 && num <= 100:
		fmt.Println("Intervalo (75,100]")
	default:
		fmt.Println("Fora de intervalo")
	}
}