package main

import (
	"fmt"
)

func sumArgs(args ...int) int {
	total := 0
	for _, arg := range args {
		total += arg
	}

	return total
}

func main() {
	fmt.Println("sum 1, 2, 3 = ", sumArgs(1, 2, 3))

}
