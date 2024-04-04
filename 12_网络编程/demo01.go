package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	fmt.Println("客户端启动")
	// 调用 Dial 函数: 参数需要指定tcp协议，需要指定服务器端的 IP + PORT
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端链接失败：err:", err)
		return
	}
	fmt.Println("链接成功，conn: ", conn)

	// 通过客户端发送单行数据，然后输出：
	reader := bufio.NewReader(os.Stdin) // os.Stdin代表终端的标准输入

	// 从终端读取一行用户输入的信息：
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("终端输入失败，err：", err)
	}

	// 将str发送到服务器端
	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("链接失败，err:", err)
	}
	fmt.Printf("终端数据通过客户端发送成功，一共发送了%d字节的数据，并退出", n)
}