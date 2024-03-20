// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow»
package main

import (
	"fmt"
	"strings"
)

func Reverse(str string) string {
	s := strings.Fields(str) // сплит по пробелам
	length := len(s)
	reverse := make([]string, length)

	for i, value := range s {
		reverse[length-1-i] = value
	}
	return strings.Join(reverse, " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println(str, Reverse(str))
}
