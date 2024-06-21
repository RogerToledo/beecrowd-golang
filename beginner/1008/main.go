package main

import "fmt"

func main() {
	var (
		num             int64
		workedHours     float64
		amountPourHours float64
	)

	fmt.Scan(&num)
	fmt.Scan(&workedHours)
	fmt.Scan(&amountPourHours)

	fmt.Printf("NUMBER = %d\n", num)

	salary := workedHours * amountPourHours

	fmt.Printf("SALARY = U$ %.2f\n", salary)
}
