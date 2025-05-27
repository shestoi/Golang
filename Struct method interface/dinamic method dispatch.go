// 18. Динамическая диспетчеризация методов
// Задача:
// Создайте интерфейс Transport с методом Move(). Реализуйте его для структур Car, Bicycle и Plane.
// Напишите функцию StartRace(transports []Transport), которая вызывает Move() для всех участников.
package main

import "fmt"

type Transport interface {
	Move()
}
type Car struct {
	//basis Transport
}
type Bicycle struct {
	//basis Transport
}
type Plane struct {
	//basis Transport
}

func (c Car) Move() {
	fmt.Println("Car Move")
}
func (b Bicycle) Move() {
	fmt.Println("Bicycle Move")
}
func (p Plane) Move() {
	fmt.Println("Plane Move")
}
func StartRace(transports []Transport) {
	for _, transport := range transports {
		transport.Move()
	}
}
func main() {
	transports := []Transport{}
	c := Car{}
	b := Bicycle{}
	p := Plane{}
	transports = append(transports, c)
	transports = append(transports, b)
	transports = append(transports, p)
	StartRace(transports)
}
