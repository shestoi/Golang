package main

import "fmt"

//5. Defer внутри цикла
//Напишите функцию с циклом, где defer вызывается на
//каждой итерации. Объясните, почему все defer выполняются после цикла.

func DeferInCycle() {
	for i := 1; i < 4; i++ {
		defer fmt.Println(i)
	}
}
func main() {
	DeferInCycle()
}
