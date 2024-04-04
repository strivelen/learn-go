package main

import (
	"fmt"
)

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	// 在管道中存放数据：
	intChan<- 10
	intChan<- 20

	// 关闭管道
	close(intChan)

	// 再次写入数据
	// intChan<- 30
	// fmt.Println(intChan)

	// 当管道关闭后，读取数据是可以的：
	num := <- intChan
	fmt.Println(num)
}