package main

import "fmt"

// // 接口的定义：定义规则、定义规范、定义某种能力
// type SayHello interface {
// 	// 声明没有实现的方法：
// 	sayHello()
// }

// // 接口的实现：定义一个结构体
// // 中国人：
// type Chinese struct {

// }
// // 实现接口的方法 --> 具体的实现：
// func (person Chinese) sayHello() {
// 	fmt.Println("你好")
// }

// // 接口的实现：定义一个结构体
// // 美国人：
// type American struct {

// }
// // 实现接口的方法 --> 具体的实现：
// func (person American) sayHello() {
// 	fmt.Println("Hello")
// }

// // 定义一个函数：专门用来各国人打招呼的函数，接收具备 SayHello 接口的能力的变量
// func greet(s SayHello) {
// 	s.sayHello()
// }

// func main() {
// 	// 创建一个中国人：
// 	c := Chinese{}
// 	// 创建一个美国人
// 	a := American{}

// 	// 美国人打招呼
// 	greet(a)
// 	// 中国人打招呼
// 	greet(c)
// }


// 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型。
// type integer int

// func (i integer) sayHello() {
//   fmt.Println("say hi + ", i)
// }

// func main() {
// 	var i integer = 10
// 	var s SayHello = i
// 	s.sayHello()
// }


// 一个自定义类型可以实现多个接口
// type AInterface interface {
//   a()
// }
// type BInterface interface {
//   b()
// }
// type Stu struct {}
// func (s Stu) a() {
//   fmt.Println("aaa")
// }
// func (s Stu) b() {
//   fmt.Println("bbb")
// }
// func main() {
//   var s Stu
//   var a AInterface = s
//   var b BInterface = s
//   a.a()
//   b.b()
// }



// //一个接口（比如A接口）可以继承多个别的接口（比如 B，C接口），这时如果要实现 A 接口，也必须将 B，C 接口的方法也全部实现。
// type CInterface interface {
//   c()
// }
// type BInterface interface {
//   b()
// }
// type AInterface interface {
//   BInterface
//   CInterface
//   a()
// }
// type Stu struct {}

// func (s Stu) a() {
//   fmt.Println("a")
// }
// func (s Stu) b() {
//   fmt.Println("b")
// }
// func (s Stu) c() {
//   fmt.Println("c")
// }

// func main() {
//   var s Stu
//   var a AInterface = s
//   a.a() // a
//   a.b() // b
//   a.c() // c
// }


// 空接口没有任何方法，所以可以裂解为所有类型都实现了空接口，也可以理解为我们可以把任何一个变量赋值给空接口。
type E interface {}
func main() {
  var num int = 10
  var e E =  num
  fmt.Println(e)
	var e2 interface{} = num
	fmt.Println(e2)
}

