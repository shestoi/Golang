// 12. Конструкторы и приватные поля
// Задача:
// Создайте структуру User с приватными полями login и password. Напишите конструктор NewUser
// (login, password string) *User и метод GetLogin() string, который возвращает логин.
package main

import "fmt"

type User struct {
	login    string
	password string
}

func (u User) GetLogin() string {
	return u.login
}
func NewUser(login string, password string) *User {
	return &User{login: login, password: password}
}

func main() {
	Dog := NewUser("Dog11", "123456")
	fmt.Println(Dog.GetLogin())
}
