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

	// 获取变量的类别：
	// 1. reType.Kind()
	k1 := reType.Kind()
	fmt.Println(k1)

	// 2. reValue.Kind()
	k2 := reValue.Kind()
	fmt.Println(k2)

	// 获取变量的类型
	// reValue 转成空接口：
	i2 := reValue.Interface()
	// 类型断言
	n, flag := i2.(Student)
	if flag == true {
		fmt.Printf("结构体的类型是：%T \n", n)
		fmt.Printf("学生的名字是：%v，年龄是：%v \n", n.Name, n.Age)
	}
	// n2 := n + 30
	// fmt.Println(n2)
}

// 定义学生结构体
type Student struct {
	Name string
	Age int
}

func main() {
	// 对基本数据类型进行反射
	// 定义一个基本数据类型：
	stu := Student{
		Name: "张三",
		Age: 18,
	}
	testReflect(stu)
}