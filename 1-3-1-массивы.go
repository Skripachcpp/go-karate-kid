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

	// массив с фиксированной длиной
	arr1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	// это уже срез и можно добавлять элементы
	// но он иногда пересоздает срез если его емкость закончилась а иногда не пересоздает
	// что то не вижу сайдэфектов, но судя по доке они должны быть
	arr2 := []int{1, 2, 3, 4, 5}
	arr3 := append(arr2, 6)
	arr4 := append(arr3, 7)
	fmt.Println("arr2 = ", arr2)
	fmt.Println("arr3 = ", arr3)
	fmt.Println("arr4 = ", arr4)
}
