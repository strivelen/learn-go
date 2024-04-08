// package main

// import (
// 	"fmt"
// 	"time"
// 	"sync"
// )

// var (
// 	testMap = make(map[int]int, 10)
// 	lock sync.Mutex
// )

// func testNum(num int) {
// 	lock.Lock()
// 	res := 1
// 	for i := 1; i <= num; i++ {
// 		res *= 1
// 	}
// 	time.Sleep(time.Second)
// 	testMap[num] = res
// 	lock.Unlock()
// }

// func main() {
// 	start := time.Now()
// 	for i := 1; i < 20; i++ {
// 		go testNum(i)
// 	}
// 	// 协程需要在 main 之后完毕
// 	// time.Sleep(time.Second * 2)
// 	lock.Lock()
// 	for key, val := range testMap {
// 		fmt.Printf("数字 %v 对应的阶乘是 %v \n", key, val)
// 	}
// 	lock.Unlock()
// 	end := time.Since(start)
// 	fmt.Println(end)
// }

package main
import (
	"fmt"
	_"strconv"
	"time"
	"sync"
	"math/rand"
)

var wg sync.WaitGroup

func writeData(intChan chan int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		v := rand.Intn(4) + 10
		intChan <- v
		fmt.Println("写入：",v, i)
	}
	close(intChan)
}

func readData(intChan chan int) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		num := <- intChan
		fmt.Println("读取：", num, i);
	}
}

func main() { // 主线程
	intChan := make(chan int, 50)

	wg.Add(2)

	go writeData(intChan)
	go readData(intChan)

	wg.Wait()
	for val := range intChan {
		fmt.Println(val)
	}
}