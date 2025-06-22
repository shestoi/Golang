// 4. Thread-safe булев флаг
// Реализуйте булев флаг, который можно безопасно устанавливать и проверять из разных горутин,
// используя только atomic.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	flag := atomic.Bool{}
	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if flag.Load() {
				fmt.Println("Flag load, goroutine :", id)
			} else {
				fmt.Println("Flag wait, goroutine :", id)
			}
		}(i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		flag.Store(true)
	}()
	wg.Wait()
}
