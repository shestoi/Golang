// 6. Атомарный генератор ID
// Реализуйте генератор уникальных идентификаторов, который гарантированно возвращает
// уникальные числа даже при вызове из множества горутин. Используйте atomic.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Глобальный счётчик
var idCounter atomic.Uint64

// Генератор уникального ID
func GenerateID() uint64 {
	return idCounter.Add(1) // атомарное увеличение и возврат нового значения
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := GenerateID()
			fmt.Println("Generated ID:", id)
		}()
	}

	wg.Wait()
}
