package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64

	notas := []float64{100, 50, 20, 10, 5, 2}
	moedas := []float64{1.00, 0.50, 0.25, 0.10, 0.05, 0.01}

	//fmt.Scan(&num)
	num = 188.91

	fmt.Println("NOTAS:")
	for _, nota := range notas {
		qtd := int(num / nota)
		num = math.Mod(num, nota)
		fmt.Printf("%d nota(s) de R$ %d.00\n", qtd, int(nota))
	}

	fmt.Println("MOEDAS:")
	for _, moeda := range moedas {
		qtd := int(num / moeda)
		num = math.Mod(num, moeda)
		fmt.Printf("%d moeda(s) de R$ %.2f\n", qtd, moeda)
	}
}
