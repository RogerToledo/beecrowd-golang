package main

import "fmt"

func main() {
	var num int

	notas := []int{100, 50, 20, 10, 5, 2, 1}

	fmt.Scan(&num)

	fmt.Println(num)

	for _, nota := range notas {
		qtd := num / nota
		num = num % nota
		fmt.Printf("%d notas(s) de R$ %d,00\n", qtd, nota)
	}
}
