# 面向对象编程 - OOP

Go语言面向对象编程说明：

> Golang 也支持面向对象编程（OOP），但是和传统的面向对象编程有区别，并不是纯粹的面向对象。所以我们说Golang支持面向对象编程特性是比较准确的。
>
> Golang没有类（class），Go语言的结构体（struct）和其它编程语言的类（class）有同等的地位，你可以理解Golang是基于struct来实现OOP特性的。
>
> Golang面向对象编程非常简洁，去掉了OOP语言的方法重载、构造函数和析构函数、隐藏的this指针等。
>
> Golang仍然有面向对象编程的继承，封装和多态的特性，只是实现方式和其它OOP语言不一样，比如继承：Golang没有extends关键字，继承是通过匿名字段来实现。

### 结构体的引入

描述一个具体的对象：

一位老师：姓名：张三   年龄：31岁   性别：女......

可以使用变量来处理：

```go
package main
import "fmt"
func main () {
  // 张三老师：姓名：张三   年龄：31岁   性别：女
  var name string = "张三"
  var age int = 31
  var sex string = "女"
  // 李四老师
  var name string = "李四"
  var age int = 28
  var sex string = "男"
}
// 缺点
// 1. 不利于数据的管理、维护
// 2. 老师的很多属性属于一个对象，用变量管理太分散了
```

### 结构体实例创建

```go
type Teacher struct {
  Name string // 变量名字大写外界可以访问这个属性
  Age int
  School string
}
func main() {
  // 创建结构体实例
  var t1 Teacher
  fmt.Println(t1) // { 0 } 这是默认值 各个属性类型的零值
  t1.Name = "张三"
  t1.Age = 20
  t1.School = "北大"
  fmt.Println(t1) // {张三 20 北大}
  fmt.Println(t1.Name) // 张三
}
func main() {
  var t Teacher = Teacher{}
  fmt.Println(t) // { 0 }
  t.Name = "李四"
  t.Age = 18
  t.School = "清华"
  fmt.Println(t) // {李四 18 清华}
}
func main() {
  var t Teacher = Teacher{"张三", 25, "国防大学"}
  fmt.Println(t) // {张三 25 国防大学}
}
func main() {
  var t *Teacher = new(Teacher) // 返回的是结构体指针
  // t是指针，t其实指向的就是地址，应该给这个地址的指向的对象的字段赋值：
  (*t).Name = "王五"
  (*t).Age = 30
  // 为了符合程序员的编程习惯，go提供了简化的赋值方式：
  t.School = "哈尔滨工业大学" // go编译器底层对t.School做了转化(*t).School
  fmt.Println(*t)
}
func main() {
  // var t *Teacher = &Teacher{"马汉", 45, "电子科大"}
  var t *Teacher = &Teacher{} // 返回的是结构体指针
  (*t).Name = "王五"
  (*t).Age = 30
  t.School = "哈尔滨工业大学" 
}
```

### 内存分析

<img src="https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240323225152692.png" alt="image-20240323225152692" style="zoom:67%;" />

### 结构体之间的转换

1. 结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段（名字、个数和类型）。

   ```go
   package main
   import "fmt"
   type Student struct {
   	Age int
   }
   type Person struct {
   	Age int
   }
   func main() {
   	var s Student = Student{10}
     var p Person = Person{10}
     s = student(p)
     fmt.Println(s)
     fmt.Println(p)
   }
   ```
   

2. 结构体进行 `type` 重新定义（相当于取别名），Golang认为是新的数据类型，但是相互可以转换。

   ```go
   package main
   import "fmt"
   type Student struct {
     Age int
   }
   type Stu Student
   func main() {
     var s1 Student = Student{19}
     var s2 Stu = Stu{19}
     s1 = Student(s2)
     fmt.Println(s1)
     fmt.Println(s2)
   }
   ```

   ### 结构体中方法的引入

   [demo01.go]: demo01.go	"demo01.go"

   1. 方法是作用在指定的数据类型上，和指定的数据类型绑定，因此自定义类型，都可以有方法，而不仅仅是`struct`。

   2. 方法的声明和调用格式。

      ```go
      type A struct {
        Num int
      }
      func (a A) test() {
        fmt.Println(a.Num)
      }
      // 调用
      var a A
      a.test() // 0 返回的是Num的零值
      ```

      

   
