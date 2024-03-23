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