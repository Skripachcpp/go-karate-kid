package main

import (
	"fmt"
	"strconv"
)

func sum(a int, b int) int {
	return a + b
}

func divide(divisible, divisor int) (int, int) {
	quotient := divisible / divisor
	remainder := divisible % divisor

	return quotient, remainder
}

func main() {
	fmt.Println("sum 1, 2 = ", sum(1, 2))

	quotient, remainder := divide(10, 3)

	fmt.Println("divide 10, 3 = ", strconv.Itoa(quotient)+",", remainder)

}
