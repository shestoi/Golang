package main

import (
	"fmt"
	"time"
)

//10. Pipeline из каналов
//Постройте конвейер (pipeline) из трёх этапов:
//Генератор чисел (1, 2, 3...).
//Умножение чисел на 2.
//Вывод результата.
//Каждый этап — отдельная goroutine, связанная каналами.
//Цель: освоить построение сложных цепочек обработки.

func generator() <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			out <- i
			time.Sleep(200 * time.Millisecond) // чтобы видеть поэтапно
		}
		close(out)
	}()
	return out
}

func multiplyByTwo(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * 2
		}
		close(out)
	}()
	return out
}

func printer(in <-chan int) {
	for result := range in {
		fmt.Println("Result:", result)
	}
}

func main() {

	nums := generator()
	doubled := multiplyByTwo(nums)
	printer(doubled)
}
