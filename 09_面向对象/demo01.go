// package main
// import "fmt"

// type A struct {
//   Num int
// }
// func (a A) test() {
//   fmt.Println(a.Num)
// }

// func main() {
// 	// 调用
// 	var a A = A{10}
// 	a.test()
// }

// // 定义 Person 结构体
// type Person struct {
// 	Name string 
// }

// // 给Person结构体绑定方法test
// func (p Person) test() {
// 	p.Name = "露露"
// 	fmt.Println(p.Name) // 露露
// }

// func main () {
// 	// 创建结构体实例
// 	var p Person
// 	p.Name = "丽丽"
// 	p.test()
// 	fmt.Println(p.Name) // 丽丽
// }



// type Person struct {
//   Name string
// }
// func (p *Person) test() {
//   (*p).Name = "露露"
//   fmt.Println(p.Name) // 露露
// }
// func main() {
//   var p Person
//   p.Name = "丽丽"
//   (&p).test()
//   fmt.Println(p.Name) // 露露
// }


package main
import "fmt"

type integer int

func (i integer) print() {
  i = 30
  fmt.Println("i = ", i) // 30
}

func (i *integer) change() {
  *i = 100
  fmt.Println("i = ", *i) // 100
}

func main() {
  var i integer = 20
  i.print()
  i.change()
  fmt.Println(i) // 100
  
}