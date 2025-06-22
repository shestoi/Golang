package main

import (
	"fmt"
	"time"
)

// 6. Таймауты в каналах
// Используйте select с time.After, чтобы:
// Ждать данных из канала не более 2 секунд.
// Если данных нет — вывести "Timeout!".
// Цель: реализовать таймауты для операций с каналами.
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(4 * time.Second)
		ch2 <- 2
	}()
	for {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
			return
		}
	}
}
