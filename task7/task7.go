// Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			mu.Lock()
			m[i] = i
			mu.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()

	for key, value := range m {
		fmt.Printf("key:\t%d\tvalue:\t%d\n", key, value)
	}
}
