package main

import "fmt"

func main() {
	var timeSpent, averageSpeed float64

	fmt.Scan(&timeSpent)
	fmt.Scan(&averageSpeed)

	dist := timeSpent * averageSpeed

	liters := dist / 12

	fmt.Printf("%.3f\n", liters)

}