package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type AtomicLimiter struct {
	maxRequests  int32
	interval     time.Duration
	requestCount int32
	stopFlag     int32
}

// Создание нового AtomicLimiter
func NewAtomicLimiter(maxRequests int32, interval time.Duration) *AtomicLimiter {
	rl := &AtomicLimiter{
		maxRequests: maxRequests,
		interval:    interval,
	}

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			if atomic.LoadInt32(&rl.stopFlag) == 1 {
				return
			}
			<-ticker.C
			atomic.StoreInt32(&rl.requestCount, 0)
		}
	}()

	return rl
}

// Проверка: можно ли разрешить запрос
func (rl *AtomicLimiter) Allow() bool {
	current := atomic.AddInt32(&rl.requestCount, 1)
	return current <= rl.maxRequests
}

// Остановка фоновой горутины
func (rl *AtomicLimiter) Stop() {
	atomic.StoreInt32(&rl.stopFlag, 1)
}

func main() {
	rl := NewAtomicLimiter(5, time.Second)
	defer rl.Stop()

	for i := 0; i < 10; i++ {
		allowed := rl.Allow()
		fmt.Printf("Request %d allowed: %v\n", i+1, allowed)
		time.Sleep(150 * time.Millisecond)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("=== After waiting ===")
	for i := 0; i < 5; i++ {
		fmt.Println("Allowed:", rl.Allow())
	}
}
