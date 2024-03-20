// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0
package main

import "fmt"

func setBit(num int64, pos uint) int64 {
	num ^= 1 << pos // XOR между значением и битом, сдвинутым на необходимую позицию
	return num
}

func main() {
	var num int64 = 128
	fmt.Println(num, setBit(num, 7))
}
