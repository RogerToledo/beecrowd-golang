package main

import "fmt"

func main() {
	var (
		numProd, qtd int64
		price, toPaid float64
	)

	for i := 0; i < 2 ; i++ {
		fmt.Scan(&numProd, &qtd, &price)
		toPaid += float64(qtd) * price 
	}
	
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", toPaid)
}