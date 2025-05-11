package main

import "fmt"

//9. Defer и возврат значения в условном блоке
//Напишите функцию, где defer изменяет возвращаемое значение, но сам return
//находится внутри if. Покажите, как defer влияет на результат.

func DeferAndReturnInBlock(a int) (result int) {
	defer func() {
		result = a + 10
	}()
	if result < 15 {
		return result
	} else {
		return result * 0
	}
}
func main() {
	fmt.Println(DeferAndReturnInBlock(10))
}
