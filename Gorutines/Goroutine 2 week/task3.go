// 3. Конфигурация с горячей перезагрузкой
// Разработайте систему управления конфигурацией, которая позволяет:
// Часто читать параметры конфигурации (много горутин)
// Редко полностью перезагружать конфигурацию (одна горутина)
// Используйте sync.RWMutex.
package main

import (
	"fmt"
	"sync"
	"time"
)

// Тип конфигурации
type Config struct {
	Settings map[string]string
}

// Менеджер конфигурации с блокировкой
type ConfigManager struct {
	mu     sync.RWMutex
	config *Config
}

// Создание нового конфигурационного менеджера
func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		config: &Config{
			Settings: map[string]string{
				"mode": "production",
			},
		},
	}
}

// Чтение конфигурации
func (cm *ConfigManager) Get(key string) (string, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	val, ok := cm.config.Settings[key]
	return val, ok
}

// Перезагрузка конфигурации (например, из файла или API)
//   - Горячая перезагрузка = обновление конфигурации без остановки приложения.
func (cm *ConfigManager) Reload(newSettings map[string]string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.config = &Config{
		Settings: newSettings,
	}
	fmt.Println("Конфигурация перезагружена")
}
func main() {
	manager := NewConfigManager()
	var wg sync.WaitGroup

	// Много горутин читают конфигурацию
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				if val, ok := manager.Get("mode"); ok {
					fmt.Printf("Горутина %d: mode = %s\n", id, val)
				}
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	// Одна горутина перезагружает конфигурацию
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond) // подождем немного
		manager.Reload(map[string]string{
			"mode": "debug",
		})
	}()

	wg.Wait()
}
