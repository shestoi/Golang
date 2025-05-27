//6. Пустой интерфейс и type assertion
//Задача: Напишите функцию PrintType, которая принимает пустой интерфейс (interface{}),
//определяет тип переданного значения и печатает его.

package main

import "fmt"

func PrintType(i interface{}) interface{} {
	switch t := i.(type) {
	case int:
		return t
	case float64:
		return t
	case string:
		return t
	case bool:
		return t
	default:
		return nil
	}
}
func main() {
	fmt.Println(PrintType(12))
	fmt.Println(PrintType("str"))
}
