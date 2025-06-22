package main

import "fmt"

// 2. Синхронизация горутин
// Напишите программу, где:
// Главная goroutine ждёт завершения двух рабочих goroutine с помощью канала.
// Рабочие goroutine отправляют сигнал (done <- true) после выполнения задачи.
// Цель: научиться синхронизировать горутины.
func main() {
	done := make(chan bool)
	go func() {
		fmt.Println("ok")
		done <- true
	}()
	go func() {
		fmt.Println("ok2")
		done <- true
	}()
	<-done
	<-done
}
