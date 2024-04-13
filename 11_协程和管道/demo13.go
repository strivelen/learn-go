package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个int管道
	intChan := make(chan int, 1)

	go func() {
		time.Sleep(time.Second * 5)
		intChan <- 10
	}()

	// 定义一个string管道
	strChan := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		strChan <- "golang"
	}()

	// fmt.Println(<-intChan) // 本身取数据就是阻塞的
	// fmt.Println(<-strChan)

	select {
	case v := <-intChan:
		fmt.Println("intChan: ", v)
	case v := <-strChan:
		fmt.Println("strChan: ", v)
		// default:
		// 	fmt.Println("防止select被阻塞") // 放开就不等了，直接走这里
	}
	fmt.Println("主线程执行完成")
}
