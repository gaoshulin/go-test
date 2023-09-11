package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex
var rwlock sync.RWMutex
var wg sync.WaitGroup
var x int64
var y int64

func main() {
	// 1、互斥锁
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	// 20
	fmt.Println(x)

	// 2、读写锁
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeLock()
	}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go readLock()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func add() {
	for i := 0; i < 10; i++ {
		// 加锁
		lock.Lock()
		x = x + 1

		// 解锁
		lock.Unlock()
	}
	wg.Done()
}

func writeLock() {
	rwlock.Lock() // 加写锁
	y = y + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒

	rwlock.Unlock() // 解写锁
	wg.Done()
}

func readLock() {
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒

	rwlock.RUnlock() // 解读锁
	wg.Done()
}
