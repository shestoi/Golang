package main

import "fmt"

//1. Базовый defer
//Напишите функцию, которая выводит "Start", затем откладывает вывод "Defer",
//а потом выводит "End". Объясните порядок вывода.

func BasicDefer() {
	fmt.Println("Start")
	defer fmt.Println("End")
}
func main() {
	BasicDefer()
}
