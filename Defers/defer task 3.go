package main

import "fmt"

//3. Defer и возвращаемые значения
//Напишите функцию, которая использует defer для изменения
//возвращаемого значения. Объясните результат.

func DeferForChangeReturn(a int) int {
	defer func() {
		a++
	}()
	return a
}
func main() {
	fmt.Println(DeferForChangeReturn(2))
}
