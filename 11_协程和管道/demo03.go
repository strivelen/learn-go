package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // 只定义无需赋值
// func main() {
// 	// 启动五个协程
// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1) // 协程开始的时候加1操作
// 		go func(n int) {
// 			fmt.Println("Hello World! +", i)
// 			wg.Done() // 协程执行完成减1操作
// 		}(i)
// 	}

// 	// 主线程一直在阻塞，什么时候wg减为0了，就停止阻塞
// 	wg.Wait()
// }
// 输出：
// Hello World! + 1
// Hello World! + 2
// Hello World! + 4
// Hello World! + 5
// Hello World! + 3

// 优化以上代码
func main() {
	// 启动五个协程
	wg.Add(5) // Add 中的数字和协程执行的次数一定要一致
	for i := 1; i <= 5; i++ {
		go func(n int) {
			defer wg.Done() // 协程执行完成减1操作
			fmt.Println("Hello World! +", i)
		}(i)
	}

	// 主线程一直在阻塞，什么时候wg减为0了，就停止阻塞
	wg.Wait()
}
