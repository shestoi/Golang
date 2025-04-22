// 4. Сравнение длины и capacity
// Напишите программу, которая:
// Создает срез с помощью make([]int, 3, 5),
// Добавляет элементы, пока len не превысит cap,
// Выводит len и cap после каждого добавления.
// Пример вывода:
// Исходный: len=3, cap=5
// После добавления 4: len=4, cap=5
// После добавления 6: len=6, cap=10 (удвоение capacity)
package main

import "fmt"

func main() {
	sl := make([]int, 3, 5)
	fmt.Println("Исходный: ", "len=", len(sl), "cap=", cap(sl))
	entCap := cap(sl)
	for i := 1; len(sl) <= entCap; i++ {
		sl = append(sl, i)
		if cap(sl) > entCap {
			fmt.Println("После добавления ", len(sl), ": len=", len(sl), " cap=", cap(sl), "(удвоение capacity)")
			break
		}
		fmt.Println("После добавления ", len(sl), ": len=", len(sl), " cap=", cap(sl))

	}
}
