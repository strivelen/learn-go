# 文件的操作

> 文件是什么？
>
> 文件时保存数据的地方，是数据源的一种，比如大家经常使用的word文档、excel文件、jpg文件...都是文件。文件最主要的作用就是保存数据，它即可以保存一张图片，也可以保存视频，声音...

os 包下的 File 结构体封装了对文件的操作：[https://pkg.go.dev/os@go1.22.1#File](https://pkg.go.dev/os@go1.22.1#File)

### 打开/关闭文件

> 示例：[demo01.go](./demo01.go)

```go
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
```

### IO的引入

> 示例：[demo02.go](./demo02.go)

![image-20240328233736059](https://raw.githubusercontent.com/strivelen/strivelen/main/learn-go/images/image-20240328233736059.png)

通过IO流对文件进行操作

##### 读取文件内容（一次性读取完成）

```go
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
```

##### 读取文件内容（带缓冲区）

> 示例： [demo03.go](./demo03.go)

读取文件的内容并显示在终端（带缓冲区的方式），适合读取比较大的文件，使用 `os.Open` 、`file.Close`、`bufio.NewReader()` 、`reader.ReadString()` 函数和方法。

> 底层有一个默认的字节长度，默认是：4096

```go
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
```

##### 写入文件

> 示例： [demo04.go](./demo04.go)

```go
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
```

##### 文件复制

> 示例： [demo05.go](./demo05.go)

```go
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
```

