// Удалить i-ый элемент из слайса
package main

import "fmt"

func removeElem(s []int, idx int) []int {
	s = append(s[:idx], s[idx+1:]...)
	return s
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = removeElem(s, 3)
	fmt.Println(s)
}
