// 2. Атомарный счетчик
// Создайте потокобезопасный счетчик, который может инкрементироваться и
// возвращать текущее значение, используя только атомарные операции (atomic).
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1) //задача для ожидания добавляется вначале каждой итерации
		go func() {
			atomic.AddInt32(&counter, 2)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&counter)) //текущее значение нашего счетчика
}
