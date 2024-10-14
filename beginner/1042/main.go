package main

import (
	"fmt"
	"sort"
)

func main() {
	var n1, n2, n3 int

	fmt.Scanln(&n1, &n2, &n3)

	numeros := []int{n1, n2, n3}

	sorted := make([]int, len(numeros))

	copy(sorted, numeros)

	sort.Ints(sorted)

	for _, s := range sorted {
		fmt.Println(s)
	}

	fmt.Println()

	for _, n := range numeros {
		fmt.Println(n)
	}
}
