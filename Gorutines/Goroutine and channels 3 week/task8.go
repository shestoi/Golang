package main

import (
	"fmt"
	"sync"
)

// 8. Канал только для чтения/записи
// Создайте функцию, которая:
// Принимает канал только для чтения (<-chan int).
// Возвращает канал только для записи (chan<- int).
// Цель: понять, как ограничивать направление каналов.

// Принимает канал только для чтения и возвращает канал только для записи

// Функция принимает <-chan int (только чтение), возвращает chan<- int (только запись)
func forReadToWrite(in <-chan int) chan<- int {
	out := make(chan int)

	go func() {
		for val := range in {
			out <- val
		}
		close(out)
	}()

	return out
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	in := make(chan int)
	out := make(chan int)

	// Пишем в канал
	writer := forReadToWrite(in)

	go func() {
		writer <- 1
		writer <- 2
		writer <- 3
		close(in)
	}()

	// Читаем и печатаем
	go func() {
		for val := range out {
			fmt.Println(val)
		}
		wg.Done()
	}()

	// Переправляем из out канала forReadToWrite в наш out
	go func() {
		for val := range in {
			out <- val
		}
		close(out)
	}()

	wg.Wait() // Ждём, пока чтение завершится
}
