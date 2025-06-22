package main

import (
	"fmt"
	"time"
)

// 5. select и мультиплексирование
// Напишите программу с двумя каналами:
// Один отправляет числа каждые 500 мс.
// Второй — каждые 1 сек.
// Используйте select, чтобы читать из того канала, который готов.
// Цель: освоить мультиплексирование каналов.
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	timer := time.NewTimer(time.Second * 5)
	go func() {
		i := 1
		for {
			ch1 <- i
			i++
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		i := 100
		for {
			ch2 <- i
			i++
			time.Sleep(time.Second)
		}
	}()
	for {
		select {
		case i := <-ch1:
			fmt.Println(i)
		case j := <-ch2:
			fmt.Println(j)
		case <-timer.C:
			fmt.Println("timer expired")
			return

		}
	}
}
