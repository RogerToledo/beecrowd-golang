package main

import "fmt"

func main() {
	var num, x int64

	for i := 0; i < 2; i++ {
		fmt.Scan(&num)
		x += num
	}

	fmt.Printf("X = %d\n", x)
}