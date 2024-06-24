package main

import "fmt"

func main() {
	var x, y float64

	fmt.Scan(&x)
	fmt.Scan(&y)

	consumption := x / y

	fmt.Printf("%.3f km/l\n", consumption)
}