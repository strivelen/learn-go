package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// 写入文件操作：
	// 打开文件：
	file, err := os.OpenFile("./write.txt", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}
	// 及时将文件关闭：
	defer file.Close()

	// 写入文件操作： => IO流 => 缓冲输出流（带缓冲区）
	writer := bufio.NewWriter(file)

	for i := 0; i < 10; i++ {
		writer.WriteString("你好，世界。\n")
	}

	// 流带缓冲区，刷新数据 => 真正写入文件中:
	writer.Flush()

	// 0666 是什么意思
	s := os.FileMode(0666).String()
	fmt.Println(s)
}