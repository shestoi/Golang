package main

import "fmt"

func UniqueArray(array []int) []int {
	seen := make(map[int]bool)
	var result []int
	for _, num := range array {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}
func main() {
	var array []int
	var input int
	fmt.Println("Введите числа для заполнения массива(чтобы прекратить заполнение введите букву или символ):")
	for {
		_, num := fmt.Scan(&input)
		if num != nil {
			break // если ввод не число — выходим из цикла
		}
		array = append(array, input)
	}
	fmt.Println(array)
	fmt.Println(UniqueArray(array))
}
