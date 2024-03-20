// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	val int
	mu  sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func (c *Counter) GetVal() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}

func main() {
	counter := Counter{}

	wg := sync.WaitGroup{}

	for i := 0; i < 777; i++ {
		wg.Add(1)
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("total", counter.GetVal())
}
