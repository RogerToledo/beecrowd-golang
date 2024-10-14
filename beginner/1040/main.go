package main

import "fmt"

func main() {
	var n1, n2, n3, n4, n5 float64

	fmt.Scanln(&n1, &n2, &n3, &n4)

	media := ((n1 * 2) + (n2 * 3) + (n3 * 4) + (n4 * 1)) / (2 + 3 + 4 + 1)

	fmt.Printf("Media: %.1f\n", media)

	switch {
	case media >= 7.0:
		fmt.Println("Aluno aprovado.")
	case media < 5.0:
		fmt.Println("Aluno reprovado.")
	default:
		fmt.Println("Aluno em exame.")
		fmt.Scan(&n5)
		fmt.Printf("Nota do exame: %.1f\n", n5)
		novaMedia := (media + n5) / 2
		
		if novaMedia >= 5.0 {
			fmt.Println("Aluno aprovado.")
		} else {
			fmt.Println("Aluno reprovado.")
		}

		fmt.Printf("Media final: %.1f\n", novaMedia)
	}
}