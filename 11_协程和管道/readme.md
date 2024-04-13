# 协程和管道

1. 什么是程序（program）

   是为完成特定任务，用某种语言编写的一组指令的集合，是一段静态代码。（程序是静态的）

2. 什么是进程（process）

   是程序的一次执行过程。正在运行的一个程序。进程作为资源分配单位，在内存中会为每个进程分配不同的内存区域。（进程是动态的）是一个动的过程，进程的生命周期：有它自身的产生、存在和消亡的过程。

3. 什么是线程（thread）

   进程可进一步细化为线程，是一个程序内部的一条执行路径。

   若一个进程同一时间并执行多个线程，就是支持多线程的。

   <img src="https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240401213733090.png" alt="image-20240401213733090" style="zoom: 67%;" />

   <img src="https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240401214315104.png" alt="image-20240401214315104" style="zoom:50%;" />

   4. 什么是协程（goroutine）

      > 示例：[demo01.go](./demo01.go)

      又称为微线程，纤程，协程是一种用户态的轻量级线程。

      作用：在执行A函数的时候，可以随时中断，去执行B函数，然后中断继续执行A函数（可以自动切换），注意这一切换过程并不是函数调用（没有调用语句），过程很像多线程，然后协程中只有一个线程在执行（协程的本质是个单线程）

​				<img src="https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240401214749408.png" alt="image-20240401214749408" style="zoom:50%;" />

> 对于单线程下，我们不可避免程序中出现io操作，但如果我们能在自己的程序中（即用户程序级别，而非操作系统级别）控制单线程下的多个任务能在一个任务遇到 io 阻塞时就将寄存器上下文和栈保存到某个其它地方，然后切换到另外一个任务去计算。在任务切回来的时候，恢复先前保存的寄存器上下文和栈，这样就保证了该线程能够最大限度地处于就绪态，即随时都可以被cpu执行的状态，相当于我们在用户程序级别将自己的io操作最大限度的隐藏起来，从而可以迷惑操作系统，让其看到：该线程好像是一直在计算，io 比较少，从而会更多的将 cpu 的执行权限分配给我们的线程（注意：线程是CPU控制的，而协程是程序自身控制的，属于程序级别的切换，操作系统完全感知不到，因而更加轻量级）

```go
package main
import (
	"fmt"
	"strconv"
	"time"
)
func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Hello Golang! +", strconv.Itoa(i))
		// 阻塞一秒
		time.Sleep(time.Second)
	}
}
func main() { // 主线程
	go test() // 开启一个协程
	for i := 1; i <= 10; i++ {
		fmt.Println("Hello World! +", strconv.Itoa(i))
		// 阻塞一秒
		time.Sleep(time.Second)
	}
}
// 注意：用 go 关键字来开启协程
```

### 主线程和协程执行流程

<img src="https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240401221251608.png" alt="image-20240401221251608" style="zoom:50%;" />

##### 主死从随：

> 主线程执行完毕后，如果协程没有执行完那么就直接死掉了。
>
> 1. 如果主线程退出了，则协程即使还没有执行完毕，也会退出
> 2. 当然协程也可以在主线程没有退出前，就自己结束了，比如完成了自己的任务。

### 启动多个协程

> 示例：[demo02.go](./demo02.go)

```go
func main() {
	// 启动多个协程
	for i := 1; i <= 5; i++ {
		go func() { // 匿名函数 + 外部变量 = 闭包
			fmt.Printf(" %v ", i) // 输出： 5  2  4  3  1 
		}()
	}
	time.Sleep(time.Second * 2)
}
// 如何修改使其正常输出
func main() {
	// 启动多个协程
	for i := 1; i <= 5; i++ {
		go func(n int) {
			fmt.Printf(" %v ", n) // 输出：2  3  5  4  1 
		}(i)
	}
	time.Sleep(time.Second * 2)
}
```

### 使用 WaitGroup 控制协程退出

> 示例：[demo03.go](./demo03.go)

WaitGroup 用于等待一组线程的结束。父现线程调用Add方法来设定应等待的线程的数量。每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用 Wait 方法阻塞至所有线程结束。=> 解决主线程在子协程结束后自动结束

```go
package main
import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup // 只定义无需赋值
func main() {
  // 如果知道要执行五个协程，那么就不需要循环中的 wg.Add(1) 了。
  wg.Add(5)
	// 启动五个协程
	for i := 1; i <= 5; i++ {
		// wg.Add(1) // 协程开始的时候加1操作
		go func(n int) {
      defer wg.Done()
			fmt.Println("Hello World! +", i)
			// wg.Done() // 协程执行完成减1操作
		}(i)
	}

	// 主线程一直在阻塞，什么时候wg减为0了，就停止阻塞
	wg.Wait()
}
// 输出：
// Hello World! + 1
// Hello World! + 2
// Hello World! + 4
// Hello World! + 5
// Hello World! + 3
```

***注意：Add 中加入的数字和协程的执行次数一定要保持一致。***

### 多个协程操作同一数据案例

> 示例：[demo04.go](./demo04.go)

```go
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
```

**注意：在理论上，这个 n 结果应该是0，无论怎么交替执行，最终想象的结果就是0，但是事实上：不是**

问题出现的原因：（图解为其中一种可能性）

<img src="https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240401231333527.png" alt="image-20240401231333527" style="zoom: 67%;" />

### 使用互斥锁同步协程

> 示例：[demo05.go](./demo05.go)
>
> 解决以上问题：有一个机制：确保：一个协程在执行逻辑的时候另外的协程不执行
>
> **锁的机制**

```go
// 加入互斥锁
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 1; i < 100000; i++ {
		// 加锁
		lock.Lock()
		n++
		// 解锁
		lock.Unlock()
	}
}
```

### 读写锁

> 示例： [demo06.go](./demo06.go)

golang中sync包实现了两种锁 `Mutex` （互斥锁）和 `RWMutex` （读写锁）

1. 互斥锁：

   其中 `Mutex` 为互斥锁，`Lock()` 加锁、`Unlock()` 解锁，使用 `Lock()` 加锁后，便不能再次对其进行加锁，直到利用 `Unlock()` 解锁对其解锁后，才能再次加锁，适用于读写不确定场景，即读写次数没有明显的区别。

   性能、效率相对来说较低

2. 读写锁

   RWMutex是一个读写锁，其经常用于读次数远远多于写次数的场景

   --- 在读的时候，数据之间不会产生影响，写和读之间才会产生影响

```go
package main
import (
	"fmt"
	"sync"
	"time"
)

// 加入读写锁
var lock sync.RWMutex

var wg sync.WaitGroup

func read(){
	defer wg.Done()
	lock.RLock() // 如果只是读数据，那么这个锁不产生影响，但是读写同时发生的时候，就会产生影响
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取成功")
	lock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	fmt.Println("开始写入数据")
	time.Sleep(time.Second)
	fmt.Println("写入成功")
	lock.Unlock()
}


func main() {
	wg.Add(6)
	// 开启协程 => 场合：读多写少
	for i := 0; i < 5; i++ {
		go read()
	}
	go write()

	wg.Wait()
}
```

### 管道 channel

> 示例：[demo07.go](./demo07.go)

**特质：**

1. 管道本质就是一个数据结构-队列
2. 数据是先进先出
3. 自身线程安全，多协程访问时，不需要加锁，`channel` 本身就是线程安全的
4. 管道有类型的，一个 `string` 的管道只能存放 `string` 类型数据

**管道的定义：`var 变量名 chan 数据类型`**

1. chan 管道关键字
2. 数据类型指的是管道的类型，里面放入数据的类型，管道是有类型的，int类型的管道只能写入整数int
3. 管道是引用类型，必须初始化才能写入数据，即 `make` 后才能使用

```go
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
	intChan<- 10
	num := 20
	intChan<- num
	intChan<- 40
	// 注意：不能存放大于容量的数据：
	// intChan<- 80

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
```

##### 管道的关闭

> 示例：[demo08.go](./demo08.go)

使用内置函数close可以关闭管道，当管道关闭后，就不能再向管道写数据了，但是仍然可以从该管道读取数据。

```go
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
```

##### 管道的遍历

> 示例：[demo09.go](./demo09.go)

管道支持 for-range 的方式进行遍历，请注意两个细节

1. 在遍历时，如果管道没有关闭，则会出现deadlock的错误
2. 在遍历时，如果管道已经关闭，则会正常遍历数据，遍历完后，就会退出遍历

```go
package main
import (
	"fmt"
)

func main() {
	var intChan chan int
	intChan = make(chan int, 100)

	// 向管道存入数据
	for i := 0; i < 100; i++ {
		intChan<- i
	}
	// 在遍历前，如果没有关闭管道，就会出现那deadlock的错误
	// 所以我们在遍历前要进行管道的关闭
	close(intChan)
	// 遍历：for-range
	for v := range intChan {
		fmt.Println("value =", v)
	}
}
```

##### 声明只读只写管道

> 示例：[demo11.go](./demo11.go)

管道可以声明为只读或只写性质

```go
package main
import (
	"fmt"
)

func main() {
	// 默认情况下，管道时双向的 => 可读可写
	// var intChan chan int
	
	// 声明为只写
	var intChan2 chan<- int // 管道具备<- 只写性质
	intChan2 = make(chan int, 3)
	intChan2<- 20
	// num := <- intChan2 // 报错
	// fmt.Println("num: ", num)
	fmt.Println("intChan2: ", intChan2)
	
	// 声明为只读
	var intChan3 <-chan int // 管道具备<- 只读性质
	if intChan3 != nil {
		num1 := <- intChan3
		fmt.Println("num1: ", num1)
	}
	// intChan3 <- 3 // 报错
}
```

**注意：管道只读只写只是管道的性质，不是管道的类型**

##### 管道的阻塞

1. 当管道只写入数据，没有读取，就会出现阻塞：

```go
package main
import (
	"fmt"
	"sync"
	_"time"
)

var wg sync.WaitGroup

// 写：
func writeData(intChan chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		intChan<- i
		fmt.Println("写入的数据为: ", i)
		// time.Sleep(time.Second)
	}

	// 管道关闭
	close(intChan)
}

// 读：
func readData(intChan chan int) {
	defer wg.Done()
	// 遍历：
	for v := range intChan {
		fmt.Println("读取的数据为：", v)
		// time.Sleep(time.Second)
	}
}

func main() { // 主线程
	// 写协程和读协程共同操作同一个管道
	intChan := make(chan int, 10)

	wg.Add(2)
	// 开启读和写的协程
	go writeData(intChan)
	// go readData(intChan)

	wg.Wait()
}
// 输出：
// fatal error: all goroutines are asleep - deadlock! // 死锁
// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0xc000028070?)
//         C:/Program Files/Go/src/runtime/sema.go:62 +0x25
// sync.(*WaitGroup).Wait(0x7d4388?)
//         C:/Program Files/Go/src/sync/waitgroup.go:116 +0x48
// main.main()
//         D:/Github/learn-go/11_协程和管道/demo12.go:42 +0x85
// exit status 2
```

2. 写的快，读的慢（管道读写频率不一致），不会出现阻塞问题

```go
package main
import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 写：
func writeData(intChan chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		intChan<- i
		fmt.Println("写入的数据为: ", i)
		// time.Sleep(time.Second)
	}

	// 管道关闭
	close(intChan)
}

// 读：
func readData(intChan chan int) {
	defer wg.Done()
	// 遍历：
	for v := range intChan {
		fmt.Println("读取的数据为：", v)
		time.Sleep(time.Second)
	}
}

func main() { // 主线程
	// 写协程和读协程共同操作同一个管道
	intChan := make(chan int, 10)

	wg.Add(2)
	// 开启读和写的协程
	go writeData(intChan)
	go readData(intChan)

	wg.Wait()
}
```

##### select 功能

> 示例：[demo13.go](./demo13.go)

用于解决多个管道的选择问题，也可以叫做多路复用，可以从多个管道中随机公开地选择一个来执行

> case后面必须进行的是io操作，不能是等值，随机去选择一个io操作。
>
> default防止select被阻塞住，加入default

```go
	select {
		case v := <- intChan :
			fmt.Println("intChan: ", v)
		case v := <- strChan :
			fmt.Println("strChan: ", v)
		default:
			fmt.Println("防止select被阻塞")
	}
```

### 协程(goroutine)和管道(channel)协同工作案例

> 示例：[demo10.go](./demo10.go)

<img src="https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240404112604620.png" alt="image-20240404112604620" style="zoom: 67%;" />

```go
package main
import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 写：
func writeData(intChan chan int) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		intChan<- i
		fmt.Println("写入的数据为: ", i)
		time.Sleep(time.Second)
	}

	// 管道关闭
	close(intChan)
}

// 读：
func readData(intChan chan int) {
	defer wg.Done()
	// 遍历：
	for v := range intChan {
		fmt.Println("读取的数据为：", v)
		time.Sleep(time.Second)
	}
}

func main() { // 主线程
	// 写协程和读协程共同操作同一个管道
	intChan := make(chan int, 50)

	wg.Add(2)
	// 开启读和写的协程
	go writeData(intChan)
	go readData(intChan)

	wg.Wait()
}
```

### defer+recover 机制处理错误

> 示例：[demo14.go](./demo14.go)

问题原因：多个协程工作，其中一个协程出现panic，导致程序崩溃

解决办法：利用 refer + recover 捕获panic进行处理，即使协程出现问题，主线程仍然不受影响可以继续执行

```go
package main
import (
	"fmt"
	"time"
)

// 输出数字：
func printNum() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 做除法操作
func devide() {
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("devide出现错误：", err)
		}
	}()
	num1 := 10
	num2 := 0 // panic: runtime error: integer divide by zero
	result := num1 / num2
	fmt.Println(result)
}

func main() {
	go printNum()
	go devide()

	time.Sleep(time.Second * 3)
}
```

