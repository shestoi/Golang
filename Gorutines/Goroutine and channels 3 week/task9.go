package main

import (
	"fmt"
	"time"
)

// 9. Остановка goroutine через канал
// Реализуйте управление goroutine:
// Создайте канал stopChan.
// Goroutine работает в бесконечном цикле, но завершается при получении сигнала из stopChan.
//
// Цель: научиться останавливать горутины безопасно.
func main() {
	//i := 0
	stopChan := make(chan struct{})
	go func() {
		i := 0
		for {
			select {
			case <-stopChan:
				fmt.Println("Channel closed, STOP")
				return
			default:
				fmt.Println(i)
				time.Sleep(time.Millisecond * 100)
				i++
			}
		}
	}()

	time.Sleep(time.Second)
	close(stopChan)
	time.Sleep(time.Millisecond * 100)
}
