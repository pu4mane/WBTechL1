// Реализовать собственную функцию sleep
package main

import "time"

func Sleep(d time.Duration) {
	<-time.After(d) //ждет указанное время и возвращает канал
}

func main() {
	Sleep(5 * time.Second)
}
