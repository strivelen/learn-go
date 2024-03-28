package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 备注：在下面的程序中不需要进行 Open/Close 操作，因为文件的打开和关闭操作被封装在ReadFile函数内部了
	// 读取文件内容
	content, err := ioutil.ReadFile("./test.txt") // 返回内容为：[]byte, err

	if err != nil {
		fmt.Println("读取文件内容出错，错误为：", err)
	}

	// 如果读取成功，将内容显示在终端
	fmt.Printf("%v", content)
	// 输出：
	// [232 191 153 230 152 175 228 184 128 228 184 170 228 190 155 230 181 139 232 175 149 231 154 132 230 150 135 228 187 182]
	fmt.Printf("%v", string(content)) // 输出：这是一个供测试的文件
}