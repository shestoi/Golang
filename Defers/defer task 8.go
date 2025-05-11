package main

import "fmt"

//8. Defer и цепочка вызовов
//Напишите функцию с несколькими defer, где один defer вызывает другую функцию,
//которая тоже содержит defer. Покажите порядок выполнения.

func DeferAndChainOfCalls() {
	defer func() {
		fmt.Println("first")
		defer func() {
			fmt.Println("second")
			defer fmt.Println("third")
			defer fmt.Println("fourth")
		}()
	}()
}
func main() {
	DeferAndChainOfCalls()
}
