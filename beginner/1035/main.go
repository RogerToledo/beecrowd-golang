package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scanln(&a, &b, &c, &d)

	sumCD := c + d
	sumAB := a + b
	evenA := a % 2
	msg := "Valores nao aceitos"

	if b > c && d > a {
		if sumCD > sumAB {
			if c >= 0 && d >= 0 && evenA == 0 {
				msg = "Valores aceitos"
			}
		}
	}
	fmt.Println(msg)
}