// Реализовать все возможные способы остановки выполнения горутины
package main

import (
	"context"
	"fmt"
	"time"
)

func ChWorker(stop <-chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("stop")
			return
		default:
			fmt.Println("...")
			time.Sleep(1 * time.Second)
		}
	}
}

func CtxWorker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop")
			return
		default:
			fmt.Println("...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	stop := make(chan bool)
	go ChWorker(stop)
	// ожидание
	time.Sleep(5 * time.Second)
	// сигнал остановки
	stop <- true
	// ожидание завершения горутины
	time.Sleep(1 * time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	go CtxWorker(ctx)
	time.Sleep(4 * time.Second)
	// сигнал остановки
	cancel()
	time.Sleep(1 * time.Second)
}
