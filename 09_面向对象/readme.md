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
     s = Student(p)
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
     *i = 100
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

### 继承

> 示例：[demo07.go](./demo07.go)

当多个结构体存在相同的属性（字段）和方法时，可以从这些结构体中抽象出结构体，在该结构体中定义这些相同的属性和方法，其它的结构体不需要重新定义这些属性和方法，只需嵌套一个匿名结构体即可。也就是说：在 Golang 中，如果一个 struct 嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性。

<img src="https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240324105801985.png" alt="image-20240324105801985" style="zoom:50%;" />

```go
// 定义动物结构体
type Animal struct {
  Age int
  Weight float32
}

// 给 Animal 绑定方法：喊叫
func (an *Animal) Shout() {
  fmt.Println("我可以大声喊叫")
}

// 给 Animal 绑定方法：自我展示
func (an *Animal) ShowInfo() {
  fmt.Printf("动物的年龄是：%v，动物的体重是：%v \n", an.Age, an.Weight)
}

// 定义结构体 Cat
type Cat struct {
  // 为了复用性，体现继承思维，加入匿名结构体: 将Animal中的字段和方法都达到复用
  Animal
}

// 对 Cat 绑定特有的方法：
func (c *Cat) scratch() {
  fmt.Println("我是小猫，我可以挠人")
}

func main() {
  // 创建Cat结构体示例
  cat := &Cat{}
  cat.Animal.Age = 3
  cat.Animal.Weight = 10.6
  cat.Animal.Shout() // 我可以大声喊叫
  cat.Animal.ShowInfo() // 动物的年龄是：3，动物的体重是：10.6 
  cat.scratch() // 我是小猫，我可以挠人
}
```

**继承的优点：提高代码的复用性、扩展性。**

##### 继承的注意事项

1. 结构体可以使用嵌套匿名结构体所有的字段和方法，即：首字母大写或者小写的字段、方法都可以使用。
2. 匿名结构体字段访问可以简化。
3. 当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，如希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分。

4. Golang 中支持多继承：如一个结构体嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而实现了多重继承。为了保证代码的简洁性，建议大家尽量不使用多重继承，很多语言就将多重继承去除了，但是Go中保留了。

5. 如嵌入的匿名结构体有相同的字段名或者方法名，则在访问时，需要通过匿名结构体类型来区分。

   ```go
   type A struct {
     a int
     b string
   }
   type B struct {
     c int
     d string
     a int
   }
   type c struct {
     A
     B
   }
   func main() {
     c := C{
       A{10, "aaa"}, 
       B{20, "ccc", 50}
     }
     fmt.Println(c.b)
     fmt.Println(c.d)
     fmt.Println(c.A.a)
     fmt.Println(c.B.a)
   }
   ```

6. 结构体的匿名字段是基本数据类型。

   ```go
   type c struct {
     A
     B
     int
   }
   c := C{
     A{10, "aaa"},
     B{20, "ccc", 50},
     888
   }
   fmt.Println(c.int) // 888
   ```

7. 嵌套匿名结构体后，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值。

   ```go
   c := C{
     A{10, "aaa"},
     B{20, "ccc", 50},
     888
   }
   ```

8. 嵌入匿名结构体的指针也是可以的。

   ```go
   type C1 struct {
     *A
     *B
     int
   }
   c1 := C1{
     &A{10, "aaa"},
     &B{20, "ccc", 50},
     888
   }
   fmt.Println(c1) // {地址 地址 888}
   fmt.Println(*c1.A) // {10 aaa}
   fmt.Println(*c1.B) // {20 ccc 50}
   ```

9. 结构体的字段可以是结构体类型的。（组合模式）

   ```go
   type D struct {
     a int
     b string
     c B // c字段类型是结构体B类型，D 和 B 是组合模式，不是继承关系
   }
   func main() {
     d := D{10, "ooo", B{66, "ppp", 999}}
     fmt.Println(d) // {10 ooo {66 ppp 999}}
     fmt.Println(d.c.d) // ppp
   }
   ```

### 接口

> 示例：[demo08.go](./demo08.go)

```go
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

// 接口的实现：定义一个结构体
// 美国人：
type American struct {

}
// 实现接口的方法 --> 具体的实现：
func (person American) sayHello() {
	fmt.Println("Hello")
}

// 定义一个函数：专门用来各国人打招呼的函数，接收具备 SayHello 接口的能力的变量
func greet(s SayHello) {
	s.sayHello()
}

func main() {
	// 创建一个中国人：
	c := Chinese{}
	// 创建一个美国人
	a := American{}

	// 美国人打招呼
	greet(a) // Hello
	// 中国人打招呼
	greet(c) // 你好
}
```

##### 总结：

1. 接口中可以定义一组方法，但不需要实现，不需要方法体。并且接口中不能包含任何变量。到某个自定义类型要使用的时候（实现接口的时候），再根据具体情况把这些方法具体实现出来。

2. 实现接口要实现 **所有的方法** 才是实现。

3. Golang 中的接口不需要显示的实现接口。Golang 中没有 `implement` 关键字。

   （Golang 中实现接口是基于方法的，不是基于接口的）

   例如：

   A 接口 a, b 方法

   B 接口 a, b 方法

   C 结构体 实现了 a, b 方法，那么C实现了A接口，也可以说实现了B接口（只要实现全部方法即可，和实际接口耦合性很低，比Java松散的多）

4. 接口的目的是为了定义规范，具体由别人来实现即可。

##### 接口注意事项

1. 接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量。

   ```go
   var s SayHello // ❌ 错误  接口本身不能创建实例
   s.sayHello()
   
   var s SayHello = c // ✔ 正确 可以指向一个实现了该接口的自定义类型的变量
   s.sayHello()
   
   ```

2. 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型。

   ```go
   type integer int
   func (i integer) sayHello() {
     fmt.Println("say hi + ", i)
   }
   // 创建实例
   var i integer = 10
   var s SayHello = i
   s.sayHello()
   ```

3. 一个自定义类型可以实现多个接口

   ```go
   type AInterface interface {
     a()
   }
   type BInterface interface {
     b()
   }
   type Stu struct {}
   func (s Stu) a() {
     fmt.Println("aaa")
   }
   func (s Stu) b() {
     fmt.Println("bbb")
   }
   func main() {
     var s Stu
     var a AInterface = s
     var b BInterface = s
     a.a()
     b.b()
   }
   ```

4. 一个接口（比如A接口）可以继承多个别的接口（比如 B，C接口），这时如果要实现 A 接口，也必须将 B，C 接口的方法也全部实现。

   ```go
   type CInterface interface {
     c()
   }
   type BInterface interface {
     b()
   }
   type AInterface interface {
     BInterface
     CInterface
     a()
   }
   type Stu struct {}
   
   func (s Stu) a() {
     fmt.Println("a")
   }
   func (s Stu) b() {
     fmt.Println("b")
   }
   func (s Stu) c() {
     fmt.Println("c")
   }
   
   func main() {
     var s Stu
     var a AInterface = s
     a.a()
     a.b()
     a.c()
   }
   ```

5. `interface` 类型默认是一个指针（引用类型），如果没有对 `interface` 初始化就使用，那么会输出 `nil`。

6. 空接口没有任何方法，所以可以裂解为所有类型都实现了空接口，也可以理解为我们可以把任何一个变量赋值给空接口。

   ```go
   type E interface {}
   func main() {
     var num int = 10
     var e E =  num
     fmt.Println(e)
   }
   ```

### 多态

变量（实例）具有多种形态。面向对象的第三大特征，在Go语言，多态特征是通过接口来实现的。可以按照统一的接口来调用不同的实现。这时接口变量就呈现不同的形态。

```go
// 在 demo08.go 中 greet 函数就实现了多态
func greet(s SayHello) { // s可以通过上下文来识别具体什么类型的实例，就体现多态
  s.sayHello()
}
```

接口体现多态特征：

1. 多态参数

   ```go
   func greet (s SayHello) {} // s 就叫多态参数
   ```

2. 多态数组

   比如：定义 `SayHello` 数组， 存放中国人结构体、美国人结构体

   ```go
   var arr [3]SayHello
   arr[0] = American{"rose"}
   arr[1] = Chinese{"张三"}
   arr[2] = Chinese{"李四"}
   fmt.Println(arr) // {{rose} {张三} {李四}}
   ```

### 断言

> 示例：[demo09.go](./demo09.go)

Go语言里面有一个语法，可以直接判断是否是该类型的变量：`value, ok = element.(T)`，这里 value 就是变量的值，ok是一个 bool 类型，element 是 interface 变量，T 是断言的类型。

```go
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
```

Type Switch 的基本用法

Type Switch 是 Go 语言中的一种特殊的 switch 语句，它比较的是类型而不是具体的值。它判断的某个接口变量的类型，然后根据具体类型再做相应处理。

```go
switch s.(type) {
  case Chinese:
		ch := s.(Chinese)
		ch.niuYangGe()
  case American:
		us := s.(American)
		us.disco()
}
```

