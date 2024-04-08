package main

import (
	"fmt"
	"time"
	"sync"
)

// 任务：
// 主线程开启一个goroutine每隔1s输出一次，10次后结束
// 在主线程中每隔2s输出一次,10次后结束
// 要求主线程和goroutine同时执行

var wg sync.WaitGroup

func print() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("协程每隔一秒打印一次")
		time.Sleep(time.Second);
	}
}

func main() {
	wg.Add(1)
	go print()
	for i := 0; i < 10; i++ {
		fmt.Println("主线程每隔两秒输出 go runtime")
		time.Sleep(time.Second)
	}
	wg.Wait()
}