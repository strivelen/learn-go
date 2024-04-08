package main

import (
	"fmt"
)

func main() {
	allChan := make(chan interface{}, 10)

	allChan <- "abc"
	allChan <- 123
	allChan <- true

	// a := <- allChan

	// fmt.Println(allChan, a, len(allChan), cap(allChan))
	close(allChan)
	// for i := 0; i < cap(allChan); i++ {
	// 	val, ok := <- allChan
	// 	if ok {
	// 		fmt.Println(i, val)
	// 	}
	// }

	str := <- allChan
	fmt.Println("单次取值：", str)

	// for {
	// 	val, ok := <- allChan
	// 	if !ok {
	// 		breakc
	// 	}
	// 	fmt.Println(val)
	// }

	for val := range allChan {
		fmt.Println(val)
	}
}