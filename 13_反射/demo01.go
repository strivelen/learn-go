package main
import (
	"fmt"
	"reflect"
)

// 利用一个函数，函数参数定义为空接口：
func testReflect(i interface{}) {
	// 调用 TypeOf函数，返回reflect.Type类型数据
	reType := reflect.TypeOf(i)
	fmt.Println("Type: ", reType)
	fmt.Printf("reType的类型是： %T \n", reType)
	// 调用 ValueOf函数，返回reflect.Value类型数据
	reValue := reflect.ValueOf(i)
	fmt.Println("Value: ", reValue)
	fmt.Printf("reValue的类型是： %T \n", reValue)

	
	// num1 := 1
	// num2 := 80 + num1
	// num2 := 80 + reValue
	// 如果真相获取reValue的数值，要调用Int()方法，返回v持有的符号整数
	num2 := 80 + reValue.Int()
	fmt.Println(num2)

	// reValue 转成空接口：
	i2 := reValue.Interface()
	// 类型断言
	n := i2.(int)
	n2 := n + 30
	fmt.Println(n2)
}

func main() {
	// 对基本数据类型进行反射
	// 定义一个基本数据类型：
	var num int = 100
	testReflect(num)
}