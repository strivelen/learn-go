package main
import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup // 只定义无需赋值
// 定义一个变量
var n int
// 加入互斥锁
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		// 加锁
		lock.Lock()
		n++
		// 解锁
		lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		n--
		lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(n) // 0
}

// 示例二：
// package main
// import (
// 	"fmt"
// 	"sync"
// )

// type Counter struct {
// 	mu		sync.Mutex
// 	count	int
// }

// func (c *Counter) add(){
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.count++
// 	fmt.Println(c.count)
// }

// func (c *Counter) sub() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.count--
// 	fmt.Println(c.count)
// }

// var wg sync.WaitGroup

// func main() {
// 	counter := Counter{}
// 	wg.Add(10)
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			defer wg.Done()
// 			counter.add()
// 			counter.sub()
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("结果：", counter.count)
// }


// 示例二：
// package main
// import (
//     "fmt"
//     "sync"
//     // "time"
// )

// type Counter struct {
//     mu    sync.Mutex
//     count int
// }

// func (c *Counter) Increment() {
//     c.mu.Lock()
//     defer c.mu.Unlock()
//     c.count++
// }

// func (c *Counter) Decrement() {
//     c.mu.Lock()
//     defer c.mu.Unlock()
//     c.count--
// }

// func (c *Counter) GetCount() int {
//     c.mu.Lock()
//     defer c.mu.Unlock()
//     return c.count
// }

// var wg sync.WaitGroup

// func main() {
//     counter := Counter{}

//     // 启动多个goroutine并发地增加和减少计数器的值
//     for i := 0; i < 10; i++ {
// 			wg.Add(1)
//         go func() {
// 						defer wg.Done()
//             for j := 0; j < 1000; j++ {
//                 counter.Increment()
//                 counter.Decrement()
//             }
//         }()
//     }

//     // 等待一段时间，以确保goroutine有足够的时间执行
//     // time.Sleep(time.Second)
// 		wg.Wait()

//     fmt.Println("Final Count:", counter.GetCount())
// }
