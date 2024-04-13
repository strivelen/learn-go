package main

import (
	"fmt"
	"sync"
	"time"
)

// 加入读写锁
var lock sync.RWMutex

var wg sync.WaitGroup

func read() {
	defer wg.Done()
	lock.RLock() // 如果只是读数据，那么这个锁不产生影响，但是读写同时发生的时候，就会产生影响
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取成功")
	lock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	fmt.Println("开始写入数据")
	time.Sleep(time.Second)
	fmt.Println("写入成功")
	lock.Unlock()
}

func main() {
	wg.Add(6)
	// 开启协程 => 场合：读多写少
	for i := 0; i < 5; i++ {
		go read()
	}
	go write()

	wg.Wait()
	fmt.Println("")
}
