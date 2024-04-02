package main
import (
	"fmt"
	"time"
)

// func main() {
// 	// 启动一个协程
// 	// 使用匿名函数，直接调用匿名函数
// 	go func(){
// 		fmt.Println(1)
// 	}()

// 	time.Sleep(time.Second * 2)
// }

// func main() {
// 	// 启动多个协程
// 	for i := 1; i <= 5; i++ {
// 		go func() { // 匿名函数 + 外部变量 = 闭包
// 			fmt.Printf(" %v ", i) // 输出：5  2  4  3  1 
// 		}()
// 	}
// 	time.Sleep(time.Second * 2)
// }

func main() {
	// 启动多个协程
	for i := 1; i <= 5; i++ {
		go func(n int) {
			fmt.Printf(" %v ", n) // 输出：2  3  5  4  1 
		}(i)
	}
	time.Sleep(time.Second * 2)
}