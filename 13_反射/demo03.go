package main
import (
	"fmt"
	"reflect"
)

func testReflect(i interface{}){
	reValue := reflect.ValueOf(i)
	// 通过SetInt()来改变值
	reValue.Elem().SetInt(40)
}

func main() {
	var num int = 100
	testReflect(&num) // 传入指针地址

	fmt.Println(num)
}