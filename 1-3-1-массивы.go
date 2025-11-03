package main

import "fmt"

func main() {
	// пустой массив
	var arr [5]int
	fmt.Println((arr))

	arr[4] = 14
	fmt.Println((arr))

	fmt.Println(len(arr))

	var matrix [4][4]int
	matrix[2][2] = 22

	fmt.Println((matrix))

	fmt.Println("matrix")
	for _, line := range matrix {
		for _, item := range line {
			fmt.Print(item, " ")
		}

		fmt.Println()
	}
}
