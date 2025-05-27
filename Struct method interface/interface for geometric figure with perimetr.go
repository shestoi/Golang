// 11. Интерфейс для геометрических фигур с периметром
// Задача:
// Расширьте интерфейс Shape из задачи 2, добавив метод Perimeter()
// float64. Реализуйте его для Rectangle и Circle.
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
	base   Shape
}
type Rectangle struct {
	width, height float64
	base          Shape
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
func main() {
	r := Rectangle{width: 10, height: 5}
	c := Circle{radius: 5}
	fmt.Println("Площадь прямоугольника :", r.Area())
	fmt.Printf("Площадь круга : %.2f\n", c.Area())
	fmt.Println("Периметр прямоугольнка :", r.Perimeter())
	fmt.Printf("Периметр круга : %.2f\n", c.Perimeter())
}
