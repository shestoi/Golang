//3. Метод с получателем по значению и указателю
//Задача: Создайте структуру Counter с полем Value. Добавьте два метода:
//
//Increment() (увеличивает Value на 1, получатель по указателю)
//
//GetValue() int (возвращает Value, получатель по значению)

package main

import "fmt"

type Counter struct {
	Value int
}

func Increment(counter *Counter) {
	counter.Value++
}
func GetValue(counter *Counter) int {
	return counter.Value
}
func main() {
	c := Counter{Value: 1}
	Increment(&c)
	fmt.Println(GetValue(&c))
}
