// 7. Ожидание горутин с sync.WaitGroup
// Запустите 4 горутины, каждая из которых увеличивает счётчик,
// и дождитесь их завершения с помощью sync.WaitGroup.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	var mu sync.Mutex
	var count int
	go func() {
		defer wg.Done() //дефер в этом месте чтобы мы точно не забыли показать что эта горутина завершена
		mu.Lock()
		count++
		fmt.Println("count", count)
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		mu.Lock()
		count++
		fmt.Println("count", count)
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		mu.Lock()
		count++
		fmt.Println("count", count)
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		mu.Lock()
		count++
		fmt.Println("count", count)
		mu.Unlock()
	}()

	wg.Wait()
}
