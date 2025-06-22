package main

import (
	"fmt"
	"time"
)

// 7. Паттерн Worker Pool
// Создайте пул из 3 рабочих goroutine.
// Главная goroutine отправляет задачи (числа) в канал, рабочие обрабатывают их (например, умножают на 2).
// Цель: разобрать популярный шаблон Worker Pool.
func worker(in <-chan int) {
	for val := range in {
		fmt.Println(val * 2)
	}

}
func main() {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go worker(ch)
	}
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	time.Sleep(time.Second)
}
