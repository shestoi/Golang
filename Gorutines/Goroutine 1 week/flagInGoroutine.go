// 10. Завершение горутин по флагу
// Создайте горутину, которая работает в бесконечном цикле,
// пока не будет установлен флаг (через атомарную переменную atomic.Bool)
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var flag atomic.Bool

	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			fmt.Println(i)
			if flag.Load() {
				break
			}
		}
	}()
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		flag.Store(true)
	}()
	wg.Wait()
}

//можно и не создавать вторую горутину, а установить флаг в основном потоке
