package main
import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 定义源文件
	filePath := "./copy.txt"
	filePath2 := "./copyed.txt"

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取有问题")
		return
	}

	// 写出文件
	err = ioutil.WriteFile(filePath2, content, 0666)
	if err != nil {
		fmt.Println("写出文件失败")
	}
}