// 5. Интерфейс Stringer (Stringer, а не fmt.Stringer)
// Задача: Реализуйте интерфейс fmt.Stringer для структуры Book (с полями Title, Author),
// чтобы при печати выводилось: "Книга: {Title}, Автор: {Author}".
package main

import "fmt"

type Stringer interface {
	String() string
}
type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Title: %s, Author: %s", b.Title, b.Author)
}
func main() {
	b := Book{
		Title:  "1984",
		Author: "George Orwell",
	}
	fmt.Println(fmt.Stringer(b)) //Можно и без fmt
}
