// 7. Кэш статистики с редкими обновлениями
// Создайте кэш статистических данных, где:
// Данные обновляются раз в минуту (одна горутина)
// Данные читаются сотни раз в секунду (много горутин)
// Примените sync.RWMutex.
// package main
//
// import (
//
//	"fmt"
//	"sync"
//
// )
//
//	type Cache struct {
//		mu   sync.RWMutex
//		data map[string]string
//	}
//
//	func NewCache() *Cache {
//		return &Cache{
//			data: make(map[string]string),
//		}
//	}
//
//	func (cache *Cache) Reload(newData map[string]string) {
//		cache.mu.Lock()
//		defer cache.mu.Unlock()
//		cache.data = newData
//	}
//
// //	func (cache *Cache) Get(key string) string {
// //		cache.mu.RLock()
// //		defer cache.mu.RUnlock()
// //		if val, ok := cache.data[key]; ok {
// //			return val
// //		}
// //		return "key not found"
// //	}
//
//	func main() {
//		wg := sync.WaitGroup{}
//		cache := NewCache()
//		cache.data["key"] = "value"
//		for i := 0; i < 10; i++ {
//			wg.Add(1)
//			go func() {
//				defer wg.Done()
//				fmt.Println(cache.data)
//			}()
//		}
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			cache.Reload(map[string]string{"key": "alhamdullilah"})
//
//		}()
//		wg.Wait()
//	}
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Тип статистики
type Stats struct {
	UsersOnline int
	CPUUsage    float64
}

// Кэш с RWMutex
type StatsCache struct {
	mu    sync.RWMutex
	stats Stats
}

// Обновление статистики (редко)
func (sc *StatsCache) Update(newStats Stats) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.stats = newStats
	fmt.Println("Статистика обновлена:", newStats)
}

// Получение статистики (часто)
func (sc *StatsCache) Get() Stats {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.stats
}
func main() {
	cache := &StatsCache{}
	var wg sync.WaitGroup

	// Горутина: обновление статистики раз в 2 секунды (вместо минуты — для демонстрации)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			newStats := Stats{
				UsersOnline: rand.Intn(1000),
				CPUUsage:    rand.Float64() * 100,
			}
			cache.Update(newStats)
			time.Sleep(1 * time.Second)
		}
	}()

	// Много читающих горутин
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				stats := cache.Get()
				fmt.Printf("Горутина %d: Статистика: %+v\n", id, stats)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
}
