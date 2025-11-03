package main

import "fmt"

func main() {
	fmt.Println("пример 1")
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\n")

	fmt.Println("пример 2")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n")

	fmt.Println("пример 3")
	j := 0
	for {
		fmt.Println("Бесконечный ц", j)

		if j > 1 {
			break
		}

		j++
	}

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

	fmt.Println("\n")
}
