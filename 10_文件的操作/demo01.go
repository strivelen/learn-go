package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./test.txt")
	
	if err != nil {
		// 当读取文件失败时，输出：
		// 文件打开出错，对应错误为:  open ./test1.txt: The system cannot find the file specified.
		fmt.Println("文件打开出错，对应错误为: ", err)
		return
	}
	// 没有出错，输出文件
	fmt.Printf("文件 = %v", file); // 文件 = &{0xc000100a08}

	// 关闭文件
	err2 := file.Close()

	if err2 != nil {
		fmt.Println("关闭文件失败：", err) // 输出：关闭文件失败： open ./test2.txt: The system cannot find the file specified.
	}
}