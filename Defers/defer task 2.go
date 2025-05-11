package main

import "fmt"

//2. Порядок выполнения defer (LIFO)
//Напишите функцию с несколькими defer и покажите, что они выполняются
//в обратном порядке (LIFO — Last In, First Out).

func DeferLifo() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
}
func main() {
	DeferLifo()
}
