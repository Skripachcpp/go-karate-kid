package main

import "fmt"

func main() {
	str := "рако мака фо фриста-а-айло"
	bytes := []byte(str)

	fmt.Println((bytes))

	// эээ, да вроде они и должны быть равны, нет?
	// string это обычно особый тип который ведет себя не как ссылочный
	fmt.Println("str == str", (string(bytes) == str))

	fmt.Println("bytes element 1", bytes[1])
	fmt.Println("str element 1", str[1])
	// fmt.Println("str element 1", string(str[1]))
}
