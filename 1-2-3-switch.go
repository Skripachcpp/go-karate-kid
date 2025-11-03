package main

import "fmt"

// Напишите программу, которая определяет название языка по его коду. Правила:

// en → English
// fr → French
// ru или rus → Russian
// иначе → Unknown
func mapLgCodeToLgName(key string) string {
	switch key {
	case "en":
		return "English"
	case "fr":
		return "French"
	case "ru":
		return "Russian"
	}

	return ""
}

func main() {
	fmt.Println(mapLgCodeToLgName("ru"))
}
