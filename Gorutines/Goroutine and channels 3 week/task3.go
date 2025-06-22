package main

import "fmt"

// 3. Буферизованный vs небуферизованный канал
// Сравните поведение:
// Небуферизованного канала (make(chan int)).
// Буферизованного (make(chan int, 3)).
// Отправьте 3 значения без чтения и посмотрите, где происходит блокировка.
// Цель: понять разницу между буферизованными и небуферизованными каналами.
func main() {
	undufferedChan := make(chan int)
	bufferedChan := make(chan int, 3)
	for i := 0; i < 3; i++ {
		select {
		case undufferedChan <- i:
			fmt.Println("unbufferedChan", i)
		case bufferedChan <- i:
			fmt.Println("bufferedChan", i)
		}
	}
}
