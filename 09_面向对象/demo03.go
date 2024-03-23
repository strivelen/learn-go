package main
import "fmt"

// 2.调用方式不一样
// type Student struct {
//   Name string
// }

// // 定义方法
// func (s Student) method01() {
// 	fmt.Println(s.Name)
// }

// // 定义函数
// func func01(s Student) {
// 	fmt.Println(s.Name)
// }

// func main() {
// 	// 创建结构体实例
// 	var s Student = Student{"张三"}
// 	// 调用方法
// 	s.method01()
// 	// 调用函数
// 	func01(s)
// }


// 3.对于函数来说，参数类型对应是什么就要传入什么
// type Student struct {
//   Name string
// }

// func func01(s Student) {
// 	fmt.Println(s.Name)
// }

// func func02(s *Student) {
// 	fmt.Println((*s).Name)
// }

// func main() {
// 	var s Student = Student{"张三"}
// 	func01(s) // 张三
// 	// func01(&s) // 错误
// 	func02(&s) // 张三
// 	// func02(s) // 错误
// }

// 4.对于方法来说，接收者为值类型，可以传入指针类型，接收者为指针类型，可以传入值类型
type Student struct {
  Name string
}

func (s Student) test01() {
	fmt.Println(s.Name)
}

func (s *Student) test02() {
	fmt.Println((*s).Name)
}

func main() {
	var s Student = Student{"丽丽"}
	s.test01()
	(&s).test01() // 虽然用指针类型调用，但是传递还是按照值传递的形式

	(&s).test02()
	s.test02() // 等价
}