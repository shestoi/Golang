package main

import "fmt"

//6. Defer и изменение именованного возвращаемого значения (в Go можно типу возвращаемого
//значения присвоить имя) Напишите функцию, где defer изменяет именованное возвращаемое
//значение после return.

func DeferAndChangingNamedReturn(a int) (result int) {
	defer func() {
		result = a + 10
	}()
	return a
}
func main() {
	fmt.Println(DeferAndChangingNamedReturn(10))
}
