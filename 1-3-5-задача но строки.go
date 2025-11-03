package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	// "strings"
	// "unicode"
)

func main() {
	// phrase := "Today I learned"
	// phraseUpper := unicode.ToUpper(phrase);
	phrase := readString()
	words := strings.Fields(phrase)

	var abbr []rune
	// var abbr []string =
	for _, word := range words {
		// wordUpper := unicode.ToUpper(word)
		runes := []rune(word)
		if len(runes) > 0 {
			firstChar := runes[0]

			if !unicode.IsLetter(firstChar) {
				continue
			}

			firstCharUpper := unicode.ToUpper(firstChar)

			abbr = append(abbr, firstCharUpper)
		}

	}

	// 1. Разбейте фразу на слова, используя `strings.Fields()`
	// 2. Возьмите первую букву каждого слова и приведите
	//    ее к верхнему регистру через `unicode.ToUpper()`
	// 3. Если слово начинается не с буквы, игнорируйте его
	//    проверяйте через `unicode.IsLetter()`
	// 4. Составьте слово из получившихся букв и запишите его
	//    в переменную `abbr`
	// ...

	fmt.Println(string(abbr))
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
