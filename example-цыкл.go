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

	fmt.Println("\n")
}
