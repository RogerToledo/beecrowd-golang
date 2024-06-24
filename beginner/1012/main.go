package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = 3.14159

	var a, b, c float64

	fmt.Scanln(&a, &b, &c)

	areaTriangulo := (a * c) / 2
	fmt.Printf("TRIANGULO: %.3f\n", areaTriangulo)

	areaCirculo := pi * math.Pow(c, 2)
	fmt.Printf("CIRCULO: %.3f\n", areaCirculo)

	areaTrapezio := ((a + b) * c) / 2
	fmt.Printf("TRAPEZIO: %.3f\n", areaTrapezio)

	areaQuadrado := b * b
	fmt.Printf("QUADRADO: %.3f\n", areaQuadrado)

	areaRetangulo := a * b
	fmt.Printf("RETANGULO: %.3f\n", areaRetangulo)

}