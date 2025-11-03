package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["key"] = 1
	m["other"] = 2

	fmt.Println("map:", m)

	m["key"] = 4
	fmt.Println("m['key']", m["key"])

	fmt.Println("m длина", len(m))

	_, existsKey := m["key"]
	fmt.Println("получаем удаленный элемент key", existsKey)

	delete(m, "key")

	fmt.Println("удалили элемент key", len(m))

	_, existsKey = m["key"]
	fmt.Println("получаем удаленный элемент key", existsKey)
}
