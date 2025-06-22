// 9. Хранилище временных меток
// Создайте потокобезопасное хранилище последних временных меток событий, где:
// Запись происходит редко (обновление метки)
// Чтение происходит часто
// Используйте sync.RWMutex.
package main

import (
	"fmt"
	"sync"
	"time"
)

// TimestampStore хранит последние временные метки событий.
type TimestampStore struct {
	mu     sync.RWMutex
	stamps map[string]time.Time
}

// NewTimestampStore создаёт и возвращает пустое хранилище.
func NewTimestampStore() *TimestampStore {
	return &TimestampStore{
		stamps: make(map[string]time.Time),
	}
}

// Set обновляет временную метку события с именем key.
func (s *TimestampStore) Set(key string, ts time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stamps[key] = ts
}

// Get возвращает последнюю временную метку события key.
// Второй возвращаемый параметр ok равен false, если метки для key нет.
func (s *TimestampStore) Get(key string) (ts time.Time, ok bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ts, ok = s.stamps[key]
	return
}

// GetAll возвращает копию всех текущих меток.
// (Полезно, если вам нужно «снять снепшот» сразу по всем событиям.)
func (s *TimestampStore) GetAll() map[string]time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Создаём копию, чтобы не выдавать внутренний map напрямую
	snapshot := make(map[string]time.Time, len(s.stamps))
	for k, v := range s.stamps {
		snapshot[k] = v
	}
	return snapshot
}
func main() {
	store := NewTimestampStore()

	// Обновляем метку
	store.Set("event-A", time.Now())

	// Много горутин могут одновременно читать:
	go func() {
		if ts, ok := store.Get("event-A"); ok {
			fmt.Println("Last A:", ts)
		}
	}()

	// Получить snapshot всех:
	snapshot := store.GetAll()
	for k, v := range snapshot {
		fmt.Printf("%s => %v\n", k, v)
	}
}
