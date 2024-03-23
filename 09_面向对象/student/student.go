package student

type student struct {
	Name string
	Age int
}

// 工厂模式
func NewStudent(n string, a int) *student {
	return &student{n, a}
}