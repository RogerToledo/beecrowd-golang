package main

import (
	"fmt"
)

func main() {
	var (
		teaType int64
		teaList = make([]int64, 5)
		count int64
	)
	
	fmt.Scan(&teaType)
	fmt.Scanln(&teaList[0], &teaList[1], &teaList[2], &teaList[3], &teaList[4])

	for _, tea := range teaList {
		if tea == teaType {
			count++
		}
	}

	fmt.Println(count)
}