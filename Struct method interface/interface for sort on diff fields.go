// 17. Интерфейс для сортировки по разным полям
// Задача:
// Создайте структуру Product (Name, Price, Rating). Реализуйте интерфейс sort.Interface для сортировки слайса
// []Product по любому из полей (выбор поля — через вложенную структуру-компаратор).
package main

type Product struct {
	Name   string
	Price  float64
	Rating float64
}
type Products []Product

func (p Products) Len() int           { return len(p) }
func (p Products) Less(i, j int) bool { return p[i].Price < p[j].Price }
func (p Products) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
