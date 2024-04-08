package main

import (
	"fmt"
	"sync"
	"runtime"
)

// 任务
// 百万级并发

var wg sync.WaitGroup

func runTime(j int) {
	defer wg.Done()
	fmt.Println("运行协程", j)
}

func main() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go runTime(i)
	}
	wg.Wait()
	// runtime.GOMAXPROCS(64) // 1.8前，要设置CPU核心数，1.8之后默认全开
	fmt.Println(runtime.NumCPU())
}