package main

import (
	"fmt"
)

func main() {
	var intChan chan int
	intChan = make(chan int, 100)

	// 向管道存入数据
	for i := 0; i < 100; i++ {
		intChan <- i
	}
	fmt.Println("遍历前 intChan 长度：", len(intChan), "容量是：", cap(intChan))
	// 在遍历前，如果没有关闭管道，就会出席那deadlock的错误
	// 所以我们在遍历前要进行管道的关闭
	close(intChan)
	// 遍历：for-range
	for v := range intChan {
		fmt.Println("value =", v)
	}
	fmt.Println("遍历完成后 intChan 长度：", len(intChan), "容量是：", cap(intChan))
}
