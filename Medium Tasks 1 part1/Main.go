package main

import (
	"fmt"
	"strconv"
)

func main() {
	first := "104"
	fmt.Printf("Тип первой переменной :%T\n", first)
	second := 35
	fmt.Printf("Тип второй переменной : %T\n", second)
	newFirst, _ := strconv.Atoi(first)
	newSecond := strconv.Itoa(second)
	fmt.Printf("Тип первой переменной после изменения :%T\n", newFirst)
	fmt.Printf("Тип второй переменной после изменения : %T", newSecond)
}
