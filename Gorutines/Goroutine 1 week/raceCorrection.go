// 6. Исправление гонки через sync.Mutex
// Модифицируйте предыдущую задачу, используя sync.Mutex, чтобы избежать гонки данных.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var i int
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(3)                         //добавляем в группу ожидания 3 горутины
	firstDone := make(chan struct{})  //создаем первый канал
	secondDone := make(chan struct{}) //создаем второй канал

	go func() {
		mutex.Lock()
		i++
		fmt.Println("first goroutine", i)
		mutex.Unlock()
		wg.Done()
		close(firstDone) //закрываем первый канал
	}()
	go func() {
		<-firstDone //передаем сигнал во второй канал, что первый завершен
		mutex.Lock()
		i++
		fmt.Println("second goroutine", i)
		mutex.Unlock()
		wg.Done()
		close(secondDone) //закрываем второй канал
	}()
	go func() {
		<-secondDone
		mutex.Lock()
		i++
		fmt.Println("third goroutine", i)
		mutex.Unlock()
		wg.Done()
	}()
	wg.Wait()
	//сделал два канала потому как вывод строки не совпадал, инкрементация аргумента производилась верно
}
