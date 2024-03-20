// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.
package main

import (
	"fmt"
)

func main() {
	num := []int{2, 4, 6, 8, 10}

	firstCh, secondCh := make(chan int), make(chan int)

	go func() {
		for _, n := range num {
			firstCh <- n
		}
		close(firstCh)
	}()

	go func() {
		for n := range firstCh {
			secondCh <- n * 2
		}
		close(secondCh)
	}()

	for n := range secondCh {
		fmt.Println(n)
	}
}
