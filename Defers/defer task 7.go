package main

import "fmt"

//7. Defer и замыкание (closure)
//Напишите функцию, где defer использует замыкание для доступа к переменной после её изменения.
//Обьяснить что такое замыкание

func DeferAndClosure(a int) int {
	defer func() {
		fmt.Println(a)
	}()
	a += 10
	return a
}
func main() {
	fmt.Println(DeferAndClosure(10))
}
