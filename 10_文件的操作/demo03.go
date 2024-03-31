package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main() {
	// 打开文件：
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("文件打开失败，err: ", err)
	}
	// 当函数退出时，让 file 关闭，防止内存泄漏
	defer file.Close()
	// 创建一个流：
	reader := bufio.NewReader(file)
	// 读取操作：
	for {
		str, err := reader.ReadString('\n') // 读取到一个换行就结束
		if err == io.EOF { // io.EOF 表示已经读取到文件的结尾
			break
		}
		// 如果没有读取到文件结尾的话，就正常输出文件内容即可：
		fmt.Println(str)
	}
	// 结束：
	fmt.Println("文件读取成功，并且全部读取完毕")
}