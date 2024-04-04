# 网络编程

把分布在不同的地理区域的计算机与专门的外部设备用通信线路互联成一个规模大、功能强的网络系统，从而使众多计算机可以方便的互相传递信息、共享硬件、软件、数据信息等资源。

设备之间在网络中进行数据的传输，发送/接收数据。

<img src="https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240404150608486.png" alt="image-20240404150608486" style="zoom:50%;" />

通信两个重要的要素：IP + PORT

设备之间进行传输的时候，必须遵照一定的规则：通信规则

<img src="https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240404182134563.png" alt="image-20240404182134563" style="zoom:67%;" />

<img src="https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240404182319475.png" alt="image-20240404182319475" style="zoom:67%;" />

### 基于TCP协议的网络通信

##### 创建客户端

> 示例：[demo01.go](./demo01.go)

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("客户端启动")
	// 调用 Dial 函数: 参数需要指定tcp协议，需要指定服务器端的 IP + PORT
	conn, err := new.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端链接失败：err:", err)
		return
	}
	fmt.Println("链接成功，conn: ", conn)
}
```

##### 创建服务器端

> 示例：[demo02.go](./demo02.go)

```go
package main

import (
	"fmt"
	"net"
)



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
	}
}
```

