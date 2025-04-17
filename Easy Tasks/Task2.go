package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Println("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Printf("Сумма ваших чисел равна : %d", a+b)
}
