package main

import "fmt"

func main() {
	var a int
	result := 1
	fmt.Println("Введите число: ")
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		result *= i
	}
	fmt.Printf("Факториалом числа %d является : %d", a, result)
}
