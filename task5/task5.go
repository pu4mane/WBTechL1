// Разработать программу, которая будет последовательно отправлять значения
// в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(666 * time.Millisecond)
		}
	}()

	go func() {
		for n := range ch {
			fmt.Println("done", n)
		}
	}()

	time.Sleep(5 * time.Second)
	close(ch)
}
