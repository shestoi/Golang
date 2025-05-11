package main

import "fmt"

//4. Defer и аргументы функции
//Напишите функцию, где defer использует переменную,
//которая изменяется после defer. Объясните результат.

func DeferAndArguments(a int) int {
	defer func() {
		fmt.Println(a)
	}()
	a++
	return a
}
func main() {
	fmt.Println(DeferAndArguments(1))
}
