// 统计 1-2000000的数字中，哪些是素数
// 传统方式：使用循环 判断
// 优化：使用并发和并行的方式，golang: 将统计分配给多个 goroutine 去完成

// ===================== 方案1 传统方式 ===================== 
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func isPrimeNum(num int) {
// 	for i := 1; i < num; i++ {
// 		var flag bool = true;
// 		for j := 2; j < i; j++ {
// 			if i % j == 0 {
// 				flag = false
// 				continue
// 			}
// 		}
// 		if flag {
// 			fmt.Println("数字", i, "是素数")
// 		}
// 	}
// }

// func main() {
// 	startTime := time.Now()
// 	isPrimeNum(20000)
// 	// 计算程序运行时间
// 	endTime := time.Now()
// 	elapsedTime := endTime.Sub(startTime)
// 	fmt.Printf("Total time taken: %v\n", elapsedTime)
// }


// ===================== 方案2 协程 ===================== 
// package main

// import (
// 	"fmt"
// 	"time"
// )

// var intChan chan int = make(chan int, 20000)

// func initChan(num int) {
// 	for i := 1; i < num; i++ {
// 		intChan <- i
// 	}
// 	close(intChan)
// }

// func isPrime(intChan chan int, primeChan chan int, exitChan chan bool) {
// 	var flag bool
// 	for {
// 		num, ok := <- intChan
// 		flag = true

// 		if !ok {
// 			break
// 		}
// 		for j := 2; j < num; j++ {
// 			if num % j == 0 {
// 				flag = false
// 				continue
// 			}
// 		}
// 		if flag {
// 			primeChan <- num
// 		}
// 	}
// 	exitChan <- true
// }

// func main() {
// 	var primeChan chan int = make(chan int, 20000)
// 	var exitChan chan bool = make(chan bool, 8)
// 	startTime := time.Now()
// 	go initChan(20000)

// 	for i := 0; i <= 8; i++ {
// 		go isPrime(intChan, primeChan, exitChan)
// 	}

// 	go func(){
// 		for i := 0; i<=8; i++ {
// 			<- exitChan
// 		}
// 		close(primeChan)
// 	}()

// 	for val := range primeChan {
// 		fmt.Println("数字", val, "是素数")
// 	}
// 	// 计算程序运行时间
// 	endTime := time.Now()
// 	elapsedTime := endTime.Sub(startTime)
// 	fmt.Printf("Total time taken: %v\n", elapsedTime)
// }


// ===================== AI方案 ===================== 
// package main

// import (
// 	"fmt"
// 	_"math"
// 	"time"
// )

// // 判断一个数是否为素数
// func isPrime(num int) bool {
// 	if num <= 1 {
// 		return false
// 	}
// 	if num <= 3 {
// 		return true
// 	}
// 	if num%2 == 0 || num%3 == 0 {
// 		return false
// 	}
// 	for i := 5; i*i <= num; i += 6 {
// 		if num%i == 0 || num%(i+2) == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	// 打印范围内的素数
// 	lower := 1
// 	upper := 20000

// 	fmt.Printf("Prime numbers between %d and %d:\n", lower, upper)
// 	// 记录程序开始时间
// 	startTime := time.Now()
// 	for num := lower; num <= upper; num++ {
// 		if isPrime(num) {
// 			fmt.Println("数字", num, "是素数")
// 		}
// 	}
// 	// 计算程序运行时间
// 	endTime := time.Now()
// 	elapsedTime := endTime.Sub(startTime)
// 	fmt.Printf("Total time taken: %v\n", elapsedTime)
// }

// ===================== AI 协程 方案 ===================== 
// package main

// import (
// 	"fmt"
// 	_"math"
// 	"sync"
// 	"time"
// )

// // 判断一个数是否为素数
// func isPrime(num int) bool {
// 	if num <= 1 {
// 		return false
// 	}
// 	if num <= 3 {
// 		return true
// 	}
// 	if num%2 == 0 || num%3 == 0 {
// 		return false
// 	}
// 	for i := 5; i*i <= num; i += 6 {
// 		if num%i == 0 || num%(i+2) == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	lower := 1
// 	upper := 20000

// 	fmt.Printf("Prime numbers between %d and %d:\n", lower, upper)

// 	// 使用 WaitGroup 来等待所有协程完成
// 	var wg sync.WaitGroup

// 	// 使用 channel 来接收素数
// 	primeCh := make(chan int)

// 	// 记录程序开始时间
// 	startTime := time.Now()

// 	// 启动多个协程进行素数判断
// 	for num := lower; num <= upper; num++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			defer wg.Done()
// 			if isPrime(n) {
// 				primeCh <- n
// 			}
// 		}(num)
// 	}

// 	// 关闭 channel，确保所有发送协程完成后可以退出循环
// 	go func() {
// 		wg.Wait()
// 		close(primeCh)
// 	}()

// 	// 从 channel 中读取并打印素数
// 	for prime := range primeCh {
// 		fmt.Println(prime)
// 	}

// 	// 计算程序运行时间
// 	endTime := time.Now()
// 	elapsedTime := endTime.Sub(startTime)
// 	fmt.Printf("Total time taken: %v\n", elapsedTime)
// }

// 这段代码由tabnine生成
package main

import (
  "fmt"
  _"math"
  "time"
	"sync"
)

func isPrime(num int) bool {
	if num <= 1 {
    return false
  }
  if num <= 3 {
    return true
  }
  if num%2 == 0 || num%3 == 0 {
    return false
  }
  for i := 5; i*i <= num; i += 6 {
    if num%i == 0 || num%(i+2) == 0 {
      return false
    }
  }
  return true
}

func main() {
	lower := 1
  upper := 20000

  fmt.Printf("Prime numbers between %d and %d:\n", lower, upper)

  // 使用 WaitGroup 来等待所有协程完成
  var wg sync.WaitGroup

  // 使用 channel 来接收素数
  primeCh := make(chan int)

  // 记录程序开始时间
  startTime := time.Now()

  // 启动多个协程进行素数判断
  for num := lower; num <= upper; num++ {
    wg.Add(1)
    go func(n int) {
      defer wg.Done()
      if isPrime(n) {
        primeCh <- n
      }
    }(num)
  }

  // 关闭 channel，确保所有发送协程完成后可以退出循环
	go func() {
    wg.Wait()
    close(primeCh)
  }()

  // 从 channel 中读取并打印素数
  for prime := range primeCh {
    fmt.Println(prime)
  }

  // 计算程序运行时间
  endTime := time.Now()
  elapsedTime := endTime.Sub(startTime)
  fmt.Printf("Total time taken: %v\n", elapsedTime)
}
