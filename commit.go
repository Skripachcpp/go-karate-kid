package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// exec выполняет команду и выводит её в консоль
func execCommand(command string) error {
	fmt.Println(command)

	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// to00 форматирует число в двузначную строку
func to00(value int) string {
	if value > 9 {
		return fmt.Sprintf("%d", value)
	}
	if value <= 0 {
		return fmt.Sprintf("0%d", -value)
	}
	return fmt.Sprintf("0%d", value)
}

func main() {
	// Получаем текущую директорию
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Ошибка получения текущей директории: %v\n", err)
		os.Exit(1)
	}

	// Меняем директорию на текущую (обычно не требуется в Go)
	err = os.Chdir(dir)
	if err != nil {
		fmt.Printf("Ошибка смены директории: %v\n", err)
		os.Exit(1)
	}

	// Получаем текущую дату
	now := time.Now()
	comment := fmt.Sprintf("%s.%s.%d",
		to00(now.Day()),
		to00(int(now.Month())),
		now.Year())

	// Выполняем git команды
	err = execCommand("git add --all")
	if err != nil {
		fmt.Printf("Ошибка при git add: %v\n", err)
		os.Exit(1)
	}

	err = execCommand(fmt.Sprintf("git commit -m '%s'", comment))
	if err != nil {
		fmt.Printf("Ошибка при git commit: %v\n", err)
		os.Exit(1)
	}

	err = execCommand(fmt.Sprintf("git push"))
	if err != nil {
		fmt.Printf("Ошибка при git push: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
