package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение: ")
	mathematical_operation, _ := reader.ReadString('\n')
	mathematical_operation = strings.TrimSpace(mathematical_operation)
	fmt.Println(isTure(mathematical_operation))
	fmt.Println(mathematical_operation)
}

func isTure(str string) bool {
	newStr := strings.Replace(str, " ", "", -1)
	for i, n := range newStr {
		if !(unicode.IsDigit(n)) && !(unicode.IsSymbol(n)) {
			return false
		} else if i != len(newStr)-1 {
			if unicode.IsSymbol(rune(newStr[i])) && unicode.IsSymbol(rune(newStr[i+1])) {
				return false
			} else if unicode.IsSymbol(rune(newStr[i])) && unicode.IsSymbol(rune(newStr[len(newStr)-1])) {
				return false
			}
		} else if i == len(newStr)-1 && unicode.IsSymbol(rune(newStr[len(newStr)-1])) {
			return false
		}
	}
	return true
}
