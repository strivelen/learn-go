package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 写：
func writeData(intChan chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		intChan <- i
		fmt.Println("写入的数据为: ", i)
		// time.Sleep(time.Second)
	}

	// 管道关闭
	close(intChan)
}

// 读：
func readData(intChan chan int) {
	defer wg.Done()
	// 遍历：
	for v := range intChan {
		fmt.Println("读取的数据为：", v)
		time.Sleep(time.Second)
	}
}

func main() { // 主线程
	// 写协程和读协程共同操作同一个管道
	intChan := make(chan int, 10)

	wg.Add(2)
	// 开启读和写的协程
	go writeData(intChan)
	go readData(intChan)

	wg.Wait()
}
