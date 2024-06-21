package main

import (
	"fmt"
)

func main() {
	var(
        name string
		fixSalary float64
        totalSales float64
     )
     
     fmt.Scanln(&name)
     fmt.Scan(&fixSalary)
     fmt.Scan(&totalSales)
     
     bonus := totalSales * 0.15

	 salary := fixSalary + bonus
     
     fmt.Printf("TOTAL = R$ %.2f\n", salary)
}