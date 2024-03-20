// Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func Swap(a, b int) (int, int) {
	a += b
	b = a - b
	a -= b
	return a, b
}

// func Swap(a, b int) (int, int) {
// 	a, b = b, a
// 	return a, b
// }

func main() {
	a, b := 4, 2
	fmt.Println("a =", a, "\nb =", b)
	a, b = Swap(a, b)
	fmt.Println("Swap:\na =", a, "\nb =", b)
}
