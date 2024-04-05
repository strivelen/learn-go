# 反射

1. 反射可以做什么？
   1. 反射可以在运行时动态获取变量的各种信息，比如变量的类型、类别等信息。
   2. 如果是结构体变量，还可以获取到结构体本身的信息（包括结构体的字段、方法）
   3. 通过反射，可以修改变量的值，可以调用关联的方法。
   4. 使用反射，需要 `import("reflect")`

2. 反射相关的函数
   1. `reflect.TypeOf(变量名)`，获取变量的类型，返回 `reflect.Type` 类型
   2. `reflect.ValueOf(变量名)` ，获取变量的值，返回 `reflect.Value` 类型（`reflect.Value`是一个结构体类型)，通过`reflect.Value`，可以获取到关于该变量的很多信息。

### 对基本数据类型进行反射

> 示例：[demo01.go](./demo01.go)

<img src="https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240404234052293.png" alt="image-20240404234052293" style="zoom:67%;" />

```go
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
```

### 对结构体类型反射

> 示例： [demo02.go](./demo02.go)

```go
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


	// reValue 转成空接口：
	i2 := reValue.Interface()
	// 类型断言
	n, flag := i2.(Student)
	if flag == true {
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
```

### 获取变量的类别

获取变量的类别，两种方式：

1. reflect.Type.Kind()
2. reflect.Value.Kind()

```go
func testReflect(i interface{}) {
  reType := reflect.TypeOf(i)
  reValue := reflect.ValueOf(i)
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
```

### 通过反射修改变量

> 示例：[demo03.go](./demo03.go)

##### 修改基本数据类型的值

```go
package main
import (
	"fmt"
	"reflect"
)

func testReflect(i interface{}){
	reValue := reflect.ValueOf(i)
	// 通过SetInt()来改变值
	reValue.Elem().SetInt(40)
}

func main() {
	var num int = 100
	testReflect(&num) // 传入指针地址

	fmt.Println(num)
}
```

##### 通过反射操作结构体的属性和方法

> 示例：[demo04.go](./demo04.go)

```go
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
}

// 给结构体绑定方法
func (s Student) CPrint() {
	fmt.Println("调用了Print()方法")
	fmt.Println("学生的名字是：", s.Name)
}

func (s Student) AGetSum(n1, n2 int) int {
	fmt.Println("调用了AGetSum()方法")
	return n1 + n2
}

func (s Student) BSet(name string, age int) {
	s.Name = name
	s.Age = age
}

func TestStudentStruct(a interface{}){
	// a 转成 reflect.Value 类型：
	val := reflect.ValueOf(a)
	fmt.Println(val)

	// 通过reflect.Value类型操作结构体内部的字段：
	n1 := val.NumField()
	fmt.Println(n1) // 2
	// 获取具体的字段
	for i := 0; i < n1; i++ {
		// 第0个字段的值是：李四第1个字段的值是：18
		fmt.Printf("第%d个字段的值是：%v \n", i, val.Field(i))
	}
	// 通过reflect.Value类型操作结构体内部的方法：
	n2 := val.NumMethod()
	fmt.Println(n2) // 3

	// 调用CPrint()方法
	// 调用方法，方法的首字母必须大写才能有对应的反射的访问权限
	// 方法的顺序按照ASCII的顺序排列的，a,b,c,,,,索引：0, 1, 2...
	val.Method(2).Call(nil)
	// 调用AGetSum方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	result := val.Method(0).Call(params)
	fmt.Println("AGetSum方法的返回值为：", result[0].Int())
}

func main() {
	s := Student{
		Name: "李四",
		Age: 18,
	}
	TestStudentStruct(s)
}
```

##### 通过反射修改结构体变量

> 示例：[demo05.go](./demo05.go)

```go
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
}

// 给结构体绑定方法
func (s Student) CPrint() {
	fmt.Println("调用了Print()方法")
	fmt.Println("学生的名字是：", s.Name)
}

func (s Student) AGetSum(n1, n2 int) int {
	fmt.Println("调用了AGetSum()方法")
	return n1 + n2
}

func (s Student) BSet(name string, age int) {
	s.Name = name
	s.Age = age
}

func TestStudentStruct(a interface{}){
	// a 转成 reflect.Value 类型：
	val := reflect.ValueOf(a)
	fmt.Println(val)

	n := val.Elem().NumField()
	fmt.Println(n)

	// 修改字段的值
	val.Elem().Field(0).SetString("张三")
}

func main() {
	s := Student{
		Name: "李四",
		Age: 18,
	}
	TestStudentStruct(&s)
	fmt.Println(s)
}
```

