package main

import "fmt"

func main() {
	var numbers []int
	var input int
	var result int
	fmt.Println("Введите числа для заполнения массива(чтобы прекратить заполнение введите букву или символ):")
	for {
		_, num := fmt.Scan(&input)
		if num != nil {
			break // если ввод не число — выходим из цикла
		}
		numbers = append(numbers, input)
		result += input
	}
	fmt.Printf("Сумма элементов вашего массива равна : %d", result)
}
