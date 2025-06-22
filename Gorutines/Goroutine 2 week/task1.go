// 1. Кэш с RW Mutex
// Реализуйте потокобезопасный кэш, где операции чтения происходят значительно чаще, чем записи.
// Используйте sync.RWMutex для оптимизации производительности.
package main

import (
	"fmt"
	"sync"
)

type Cash struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewCash() *Cash {
	return &Cash{
		data: make(map[string]string),
	}
}

// Запись в кэш (только один поток одновременно)
func (c *Cash) Add(key, value string) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

// Чтение из кэша (быстрое, с RLock)
func (c *Cash) Get(key string) (string, bool) {
	c.mu.RLock()         // несколько горутин могут читать одновременно
	defer c.mu.RUnlock() // освобождаем блокировку

	val, ok := c.data[key]
	return val, ok
}
func main() {
	cache := NewCash()
	var wg sync.WaitGroup

	// Пишем одно значение
	wg.Add(1)
	go func() {
		defer wg.Done()
		cache.Add("foo", "bar")
	}()

	// Много читаем
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			val, ok := cache.Get("foo")
			if ok {
				fmt.Printf("Горутина %d: прочитала значение: %s\n", id, val)
			} else {
				fmt.Printf("Горутина %d: ключа нет\n", id)
			}
		}(i)
	}

	wg.Wait()
}
