package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Создаем контекст и отменяем его при получении сигнала прерывания
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// канал для передачи данных
	ch := make(chan string)

	// горутина для записи данных в канал
	go func() {
		for i := 0; ; i++ {
			// Проверяем, не был ли контекст отменен
			select {
			case <-ctx.Done():
				return
			default:
				// Записываем данные в канал
				ch <- fmt.Sprintf("Data %d", i)
			}
		}
	}()

	// воркеры для чтения данных из канала и вывода их в stdout
	numWorkers := 5
	wg := sync.WaitGroup{}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case data := <-ch:
					// Выводим данные в stdout
					fmt.Println(data)
				}
			}
		}()
	}

	// сигнала прерывания
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	// Отменяем контекст, чтобы завершить работу всех горутин
	cancel()

	// Ждем завершения работы всех воркеров
	wg.Wait()
}
