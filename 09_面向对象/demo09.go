package main

import "fmt"

// 接口的定义：定义规则、定义规范、定义某种能力
type SayHello interface {
	// 声明没有实现的方法：
	sayHello()
}

// 接口的实现：定义一个结构体
// 中国人：
type Chinese struct {

}
// 实现接口的方法 --> 具体的实现：
func (person Chinese) sayHello() {
	fmt.Println("你好")
}
func (person Chinese) niuYangGe() {
	fmt.Println("东北文化-扭秧歌")
}

// 接口的实现：定义一个结构体
// 美国人：
type American struct {

}
// 实现接口的方法 --> 具体的实现：
func (person American) sayHello() {
	fmt.Println("Hello")
}
func (person American) disco() {
	fmt.Println("disco")
}

// 定义一个函数：专门用来各国人打招呼的函数，接收具备 SayHello 接口的能力的变量
func greet(s SayHello) {
	s.sayHello()
	// s.niuYangGe() // 报错：s.niuYangGe undefined (type SayHello has no field or method niuYangGe)
	// 断言：
	// ch, flag := s.(Chinese) // 判断s是否是Chinese类型并赋值给ch变量,flag代表是否是Chinese类型的boolean
	// if flag {
	// 	ch.niuYangGe()
	// } else {
	// 	fmt.Println("美国人不会扭秧歌")
	// }

	// 简洁语法
	// if ch, flag := s.(Chinese); flag {
	// 	ch.niuYangGe()
	// } else {
	// 	fmt.Println("美国人不会扭秧歌")
	// }

	switch s.(type) {
		case Chinese:
			ch := s.(Chinese)
			ch.niuYangGe()
		case American:
			us := s.(American)
			us.disco()
	}
	fmt.Println("打招呼...")
}

func main() {
	// 创建一个中国人：
	c := Chinese{}
	// 中国人打招呼
	greet(c)


	// 创建一个美国人
	// a := American{}
	// 美国人打招呼
	// greet(a)
	
}