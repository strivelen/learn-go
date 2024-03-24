package main
import "fmt"

// // 定义动物结构体
// type Animal struct {
//   Age int
//   weight float32
// }

// // 给 Animal 绑定方法：喊叫
// func (an *Animal) Shout() {
//   fmt.Println("我可以大声喊叫")
// }

// // 给 Animal 绑定方法：自我展示
// func (an *Animal) showInfo() {
//   fmt.Printf("动物的年龄是：%v，动物的体重是：%v \n", an.Age, an.weight)
// }

// // 定义结构体 Cat
// type Cat struct {
//   // 为了复用性，体现继承思维，加入匿名结构体
//   Animal
// }

// // 对 Cat 绑定特有的方法：
// func (c *Cat) scratch() {
//   fmt.Println("我是小猫，我可以挠人")
// }

// func main() {
//   // 创建Cat结构体示例
//   cat := &Cat{}
//   cat.Animal.Age = 3
//   cat.Animal.weight = 10.6
//   cat.Animal.Shout()
//   cat.Animal.showInfo()
//   cat.scratch()
// }


type A struct {
  a int
  b string
}
type B struct {
  c int
  d string
  a int
}
type C struct {
  A
  B
	int
}

type D struct {
  a int
  b string
  c B // c字段类型是结构体B类型，D 和 B 是组合模式，不是继承关系
}

func main() {
  c := C{
		A{10, "aaa"}, 
		B{20, "ccc", 50}, 
		888,
	}
  fmt.Println(c.b) // aaa
  fmt.Println(c.d) // ccc
  fmt.Println(c.A.a) // 10
  fmt.Println(c.B.a) // 50
	fmt.Println(c.int) // 888


	d := D{10, "ooo", B{66, "ppp", 999}}
  fmt.Println(d) // {10 ooo {66 ppp 999}}
  fmt.Println(d.c.d) // ppp
}