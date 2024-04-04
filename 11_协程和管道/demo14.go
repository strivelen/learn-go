package main
import (
	"fmt"
	_"time"
	"sync"
)

var wg sync.WaitGroup

// 输出数字：
func printNum() {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 做除法操作
func devide() {
	defer func(){
		wg.Done()
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
	wg.Add(2)
	go printNum()
	go devide()

	// time.Sleep(time.Second * 3)
	wg.Wait()
}