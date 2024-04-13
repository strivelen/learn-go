package main

import (
	"fmt"
)

func main() {
	// 定义管道、声明管道 => 定义一个int类型的管道
	var intChan chan int
	// 通过 make 初始化：管道可以存放3个int类型的管道
	intChan = make(chan int, 3)

	// 证明管道是引用类型
	fmt.Printf("intChan 的值：%v \n", intChan) // intChan 的值：0xc00010a080

	// 向管道存放数据：
	intChan <- 10
	num := 20
	intChan <- num
	intChan <- 40
	// 注意：不能存放大于容量的数据：
	// intChan <- 80

	// 在管道中读取数据：
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Println(num1, num2, num3)

	// 注意：在没有使用协程的情况下，如果管道的数据已经全部取出，那么再取就会报错：
	// num4 := <-intChan
	// fmt.Println(num4)

	// 输出管道的长度：
	fmt.Printf("管道的实际长度：%v，管道的容量是：%v \n", len(intChan), cap(intChan))
}
