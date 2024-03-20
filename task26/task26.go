// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false
package main

import (
	"fmt"
	"strings"
)

func Unique(str string) bool {
	s := make(map[rune]bool)
	str = strings.ToLower(str)

	for _, ch := range str {
		if s[ch] {
			return false
		}
		s[ch] = true
	}

	return true
}

func main() {
	fmt.Println(Unique("abcd"), Unique("abCdefAaf"), Unique("aabcd"))
}
