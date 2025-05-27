// 8. Интерфейс для работы с БД
// Задача: Создайте интерфейс Database с методами Save(data string) и Read() string.
// Реализуйте его в структуре MockDB, которая сохраняет данные в памяти (в слайс).
package main

import "fmt"

type Database interface {
	Save(data string)
	Read() string
}

func (s *MockDB) Save(data string) {
	s.slice = append(s.slice, data)
}
func (s *MockDB) Read() string { //С его помощью выводим последний добавоенный элемент
	if len(s.slice) == 0 {
		return "База данных пуста"
	} else {
		return s.slice[len(s.slice)-1] //выводить все хранилище
	}
}

type MockDB struct {
	slice []string
}

func main() {
	db := MockDB{}
	db.Save("hi")
	//PrintType(a)
	fmt.Println(db.Read())
	db.Save("bye")
	fmt.Println(db.Read())
	fmt.Println(db)
}
