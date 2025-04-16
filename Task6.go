package main

import (
	"fmt"
	"math"
)

func main() {
	var numbers []int
	var input int
	max := math.MinInt32
	fmt.Println("Введите числа для заполнения массива(чтобы прекратить заполнение введите букву или символ):")
	for {
		_, num := fmt.Scan(&input)
		if num != nil {
			break // если ввод не число — выходим из цикла
		}
		numbers = append(numbers, input)
	}
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	fmt.Println("Максимальное значение в вашем массиве: ", max)
}
