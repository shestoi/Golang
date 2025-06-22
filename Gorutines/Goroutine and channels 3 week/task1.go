package main

import (
	"fmt"
)

// 1. Простой обмен сообщениями
// Создайте два goroutine:
// Один отправляет числа от 1 до 5 в канал.
// Второй читает их и выводит в консоль.
// Цель: понять базовую работу каналов.
func main() {

	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println("Получено:", num)
	}
}
