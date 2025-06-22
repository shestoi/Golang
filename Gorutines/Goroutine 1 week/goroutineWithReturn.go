// 8. Горутины и возврат значений
// Создайте функцию, которая запускает горутину, вычисляющую квадрат числа,
// и возвращает результат через разделяемую переменную (с синхронизацией).
// Разделяемая переменная — это переменная, доступ к которой одновременно могут иметь несколько горутин
// (то есть параллельно выполняющихся функций).
package main

import (
	"fmt"
	"sync"
)

func Square(i int, res *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	*res = i * i
	mu.Unlock()
	//Mutex так как задача подразумевает синхранизацию
}
func main() {
	var (
		wg  sync.WaitGroup
		i   int
		mu  sync.Mutex
		res int
	)
	fmt.Scanln(&i)
	wg.Add(1)
	go Square(i, &res, &wg, &mu)
	wg.Wait()
	fmt.Println(res)
}
