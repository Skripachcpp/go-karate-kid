package main

import (
	"fmt"
	"time"
)

func diff(from string, to string) float64 {
	fromDuration, _ := time.ParseDuration(from)
	fromMin := time.Duration.Minutes(fromDuration)

	toDuration, _ := time.ParseDuration(to)
	toMin := time.Duration.Minutes(toDuration)

	res := toMin - fromMin

	if res < 0 {
		return res * -1
	} else {
		return res
	}
}

func main() {
	fmt.Println("привет мир")         // вывести строку
	fmt.Println("1+1 =", 1+1)         // 1+1 = 2
	fmt.Println("5.0/2.0 =", 5.0/2.0) // 5.0/2.0 = 2.5
	fmt.Println(true && false)        // false
	fmt.Println(true || false)        // true
	fmt.Println(!true)                // false
	fmt.Println("go" + "lang")        // golang

	// Напишите программу, которая считает количество минут во временном отрезке.
	fmt.Println("1h30m = 90 min", diff("1h30m", "90 min"))
	fmt.Println("300s = 5 min", diff("300s", "5 min"))
	fmt.Println("10m = 10 min", diff("10m", "10 min"))
}
