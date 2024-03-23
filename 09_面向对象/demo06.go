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