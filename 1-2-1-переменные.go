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
	var b bool = true
	fmt.Println(b)
	// true

	var s string = "hello"
	fmt.Println(s)
	// hello

	var i int = 42
	fmt.Println(i)
	// 42

	var f float64 = 12.34
	f = 10.0
	fmt.Println(f)
	// 12.34

	const c float64 = 12.34
	// c = 1
	fmt.Println(f)
	// 12.34
}
