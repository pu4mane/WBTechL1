// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode
package main

import (
	"fmt"
)

func Reverse(str string) string {
	runes := []rune(str)
	length := len(runes)
	reversed := make([]rune, length)
	for i, s := range runes {
		reversed[length-1-i] = s
	}
	return string(reversed)
}

func main() {
	str := "главрыба"
	fmt.Println(str, Reverse(str))
}
