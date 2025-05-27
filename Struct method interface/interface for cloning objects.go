// 15. Интерфейс для клонирования объектов
// Задача:
// Создайте интерфейс Cloner с методом Clone() Cloner. Реализуйте его для структуры Document
// (поле Content string), чтобы метод возвращал копию объекта.
package main

import "fmt"

type Cloner interface {
	Clone() Cloner
}
type Document struct {
	Content string
}

func (d Document) Clone() Cloner {
	//d.Content = d.Content[:len(d.Content)-6]
	return Document{Content: d.Content}
}
func main() {
	d := Document{Content: "Hello World"}
	fmt.Println(d.Clone())
	//fmt.Println(d)
}
