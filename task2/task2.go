// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
package main

import (
	"fmt"
	"sync"
)

func main() {
	num := []int64{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	wg.Add(len(num)) //cчетчик

	for _, n := range num {
		go func(n int64) {
			fmt.Println(n * n)
			wg.Done() //-1
		}(n)
	}

	wg.Wait() //блок main wg != 0
}

// func main() {
// 	num := []int64{2, 4, 6, 8, 10}

// 	ch := make(chan int64, len(num))

// 	for _, n := range num {
// 		go func(n int64) {
// 			fmt.Println(n * n)
// 			ch <- n
// 		}(n)
// 	}

// 	for range num {
// 		<-ch
// 	}
// }
