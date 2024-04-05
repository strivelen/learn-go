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