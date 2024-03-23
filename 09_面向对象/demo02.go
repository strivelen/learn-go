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