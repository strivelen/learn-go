package main
import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup // 只定义无需赋值
// 定义一个变量
var n int

func add() {
	defer wg.Done()
	for i := 1; i < 100000; i++ {
		n = i + 1
	}
}

func sub() {
	defer wg.Done()
	for i := 1; i < 100000; i++ {
		n = i - 1
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(n) // 99998
}