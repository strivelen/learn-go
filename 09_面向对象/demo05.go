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