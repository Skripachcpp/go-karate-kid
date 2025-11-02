package main

import (
	"fmt"
	"strings"
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
	// считываем временной отрезок из os.Stdin
	// гарантируется, что значение корректное
	// не меняйте этот блок
	var s string
	// fmt.Scanln(&s)
	s = "12m=11m"

	// d, _ := time.ParseDuration(s)

	// выведите исходное значение (s)
	// и количество минут в нем
	// в формате "исходное = X min"
	// используйте метод .Minutes() объекта d

	parts := strings.Split(s, "=")
	if len(parts) <= 1 {
		return
	}

	fmt.Println(diff(parts[0], parts[1]))
}
