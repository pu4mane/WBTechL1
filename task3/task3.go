// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….)
// с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"sync"
)

func main() {
	num := []int{2, 4, 6, 8, 10}
	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(len(num))

	for _, n := range num {
		go func(n int) {
			ch <- n * n
			wg.Done()
		}(n)
	}

	//предотвращение дедлока
	go func() {
		wg.Wait()
		close(ch)
	}()

	var sum int
	for sq := range ch {
		sum += sq
	}

	fmt.Println(sum)
}
