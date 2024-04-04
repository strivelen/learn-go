package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 连接用完一定要关闭：
	defer conn.Close()

	for {
		// 创建一个切片，将读取的数据放入切片
		buf := make([]byte, 1024)

		// 从conn连接中读取数据
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		// 将读取的内容在服务器输出：
		fmt.Println(string(buf[0:n]))
	}
}

func main() {
	fmt.Println("服务器端启动了")
	// 进行监听：需要指定服务器端TCP协议，服务器端的IP+PORT
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("监听失败，err:", err)
		return
	}
	// 监听成功以后：
	// 循环等待客户端的链接
	for {
		conn, err2 := listen.Accept()
		if err2 != nil {// 等待客户端失败
			fmt.Println("客户端的等待失败，err2: ", err2)
		} else {
			// 链接成功
			fmt.Printf("等待链接成功，conn=%v, 接收到的客户端信息：%v \n", conn, conn.RemoteAddr().String())
		}

		// 准备一个协程，协程处理客户端服务请求：
		go process(conn) // 不同的客户端的请求，连接conn不一样
	}
}