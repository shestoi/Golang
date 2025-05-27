// 9. Полиморфизм с интерфейсами
// Задача: Создайте интерфейс Animal с методом Sound() string. Реализуйте его для структур Dog и Cat.
// Затем создайте функцию MakeSound(animal Animal), которая вызывает Sound().
package main

import "fmt"

type Animal interface {
	Sound() string
}
type Dog struct {
	//basis Animal
}
type Cat struct {
	//basis Animal
}

func (d Dog) Sound() string {
	return "Гав-Гав"
}
func (c Cat) Sound() string {
	return "Мяу-Мяу"
}
func MakeSound(animal Animal) {
	fmt.Println(animal.Sound())
}
func main() {
	d := Dog{}
	c := Cat{}
	MakeSound(d)
	MakeSound(c)
	fmt.Println(d.Sound())
	fmt.Println(c.Sound())
}
