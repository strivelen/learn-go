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

> 示例： [demo01.go](./demo01.go)

1. 方法是作用在指定的数据类型上，和指定的数据类型绑定，因此自定义类型，都可以有方法，而不仅仅是`struct`。
2. 方法的声明和调用格式

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

```go
// 定义 Person 结构体
type Person struct {
	Name string 
}

// 给Person结构体绑定方法test
func (p Person) test() {
	p.Name = "露露"
	fmt.Println(p.Name) // 露露
}

func main () {
	// 创建结构体实例
	var p Person
	p.Name = "丽丽"
	p.test()
	fmt.Println(p.Name) // 丽丽
}
```

注意：

1. `test`  方法中参数名字随意起
2. 结构体 `Person` 和 `test` 方法绑定，调用 `test` 方法必须靠指定的类型：`Person`
3. 如果其它类型变量调用 `test` 方法一定会报错
4. 结构体对象传入 `test` 方法中，值传递（和函数参数传递一致）

##### 方法的注意事项

1. 结构体类型是值类型，在方法调用中，遵守值类型的传递机制，是值拷贝方式

   <img src="https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240323235521493.png" alt="image-20240323235521493" style="zoom:50%;" />

2. 如果程序员希望在方法中改变结构体变量的值，可以通过结构体指针的方式处理

   ```go
   package main
   import "fmt"
   
   type Person struct {
     Name string
   }
   func (p *Person) test() {
     (*p).Name = "露露"
     fmt.Println(p.Name) // 露露
   }
   func main() {
     var p Person
     p.Name = "丽丽"
     (&p).test()
     fmt.Println(p.Name) // 露露
   }
   ```

   <img src="https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240324000546432.png" alt="image-20240324000546432" style="zoom:33%;" />

3. Golang 中的方法作用在指定的数据类型上的，和指定的数据类型绑定，因此自定义类型，都可以有方法，而不仅仅是 `struct` ，比如：`int`、`float32` 等都可以有方法。

   ```go
   package main
   import "fmt"
   
   type integer int
   
   func (i integer) print() {
     i = 30
     fmt.Println("i = ", i)
   }
   
   func (i *integer) change() {
     *i = 30
     fmt.Println("i = ", i) // 100
   }
   
   func main() {
     var i integer = 20
     // i.print()
     i.change()
     fmt.Println(i) // 100  
   }
   ```

4. 方法的访问范围控制的规则和函数一样：方法名首字母小写，只能在本包访问，方法首字母大写，可以在本包和其他包访问。

5. 如果一个类实现了 `String()` 这个方法，那么 `fmt.Println()` 默认会调用这个变量的 `String()` 进行输出。

   > 示例：[demo02.go](./demo02.go)

   ```go
   package main
   import "fmt"
   
   type Student struct {
     Name string
     Age int
   }
   
   func (s *Student) String() string {
     str := fmt.Sprintf("Name = %v, Age = %v", s.Name, s.Age)
     return str
   }
   
   func main() {
     stu := Student{
       Name: "李四",
       Age: 20,
     }
     // 传入地址，如果绑定了String方法就会自动调用
     fmt.Println(&stu)
   }
   ```

### 方法和函数的区别

> 示例：[demo03.go](./demo03.go)

1. 绑定指定类型

   - 方法：需要绑定指定数据类型
   - 函数：不需要绑定数据类型

2. 调用方式不一样

   函数的调用方式：函数名(实参列表)

   方法的调用方式：变量.方法名(实参列表)

   ```go
   type Student struct {
     Name string
   }
   
   // 定义方法
   func (s Student) method01() {
   	fmt.Println(s.Name)
   }
   
   // 定义函数
   func func01(s Student) {
   	fmt.Println(s.Name)
   }
   
   func main() {
   	// 创建结构体实例
   	var s Student = Student{"张三"}
   	// 调用方法
   	s.method01()
   	// 调用函数
   	func01(s)
   }
   ```

3. 对于函数来说，参数类型对应是什么就要传入什么

   ```go
   type Student struct {
     Name string
   }
   
   func func01(s Student) {
   	fmt.Println(s.Name)
   }
   
   func func02(s *Student) {
   	fmt.Println((*s).Name)
   }
   
   func main() {
   	var s Student = Student{"张三"}
   	func01(s) // 张三
   	// func01(&s) // 错误
   	func02(&s) // 张三
   	// func02(s) // 错误
   }
   ```

4. 对于方法来说，接收者为值类型，可以传入指针类型，接收者为指针类型，可以传入值类型

   ```go
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
   ```

### 创建结构体实例时指定字段值

> 示例：[demo04.go](./demo04.go)

```go
type Student struct {
  Name string
  Age int
}
func main() {
  // 方式1：按照顺序赋值操作
  var s1 Student = Student{"小李", 19} // 缺点：必须按照顺序有局限性
  fmt.Println(s1)
  
  // 方式2：按照指定类型
  var s2 Student = Student{
    Name: "张三",
    Age: 20,
  }
  fmt.Println(s2)
  
  // 方式3：想要返回结构体的指针类型
  var s3 *Student = &Student{"明明", 18}
  fmt.Println(*s3)
  var s4 *Student = &Student{
    Name: "李四",
    Age: 29
  }
  fmt.Println(*s4)
}
```

### 跨包创建结构体实例

> 示例：[demo05.go](./demo05.go)

> 当结构体首字母小写，如果做到跨包访问？-- **工厂模式**

```go
// ./student/student.go
package student

type student struct {
	Name string
	Age int
}

// 工厂模式
func NewStudent(n string, a int) *student {
	return &student{n, a}
}

// ================================================

// ./demo05.go
package main

import (
	"fmt"
	"example.com/student"
)

func main() {
	s := student.NewStudent("张三", 20)
	fmt.Println(*s) // {张三 20}
	fmt.Println(s) //&{张三 20}
}
```

### 封装

> 示例：[demo06.go](./demo06.go)

1. 什么是封装？

   封装（encapsulation）就是把抽象出的字段和对字段的操作封装在一起，数据被保护在内部，程序的其它包只有通过被授权的操作方法，才能对字段进行操作。

2. 封装的好处：

   - 隐藏实现细节
   - 提可以对数据进行验证，保证安全合理

3. Golang 中如何实现封装：

   - 建议将结构体、字段（属性）的首字母小写（其它包不能使用，类似 private，实际开发不小写也可能，因为封装没有那么严格）

   - 给结构体所在包提供一个工厂模式的函数，首字母大写（类似一个构造函数）

   - 提供一个首字母大写的 set 方法（类似其它语言的 public ），用于对属性判断并赋值

     func (var 结构体类型名)SetXxx（参数里欸包）{

     ​	// 加入数据验证的业务逻辑

     ​	var.Age = 参数

     }

   - 提供一个首字母大写的Get方法（类似其它语言的public），用于获取属性值

     func (var 结构体类型名) GetXxx() （返回值列表）{

     ​	return var.字段

     }

   ```go
   // person.go
   package person
   import "fmt"
   
   type person struct {
   	Name string
   	age int // 其他包不能直接访问
   }
   
   // 定义工厂模式的函数，相当于构造器
   func NewPerson(name string) *person {
   	return &person{
   		Name: name,
   	}
   }
   
   // 定义set和get函数，对age字段进行封装，因为在函数中可以加一系列的限制操作，确保被封装字段的安全合理性
   func (p *person) SetAge(age int) {
   	if age > 0 && age < 150 {
   		p.age = age
   	} else {
   		fmt.Println("对不起，你传入的年龄范围不正确")
   	}
   }
   
   func (p *person) GetAge() int {
   	return p.age
   }
   
   // ====================================================
   // demo06.go
   package main
   
   import (
   	"fmt"
   	"example.com/person"
   )
   
   func main() {
   	p := person.NewPerson("李四")
   	p.SetAge(180)
   
   	fmt.Println(p.Name) // 李四
   	fmt.Println(p.GetAge()) // 18
   	fmt.Println(*p) // {李四 18}
   }
   ```

   

   

   

​	

















