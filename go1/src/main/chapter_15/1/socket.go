package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	var (
		host          = "www.google.com"
		port          = "8888"
		remote        = host + ":" + port
		msg    string = "GET / \n"
		data          = make([]uint8, 4096)
		read          = true
		count         = 0
	)
	// 创建一个 socket
	con, err := net.Dial("tcp", remote)
	// 发送我们的消息，一个 http GET 请求
	io.WriteString(con, msg)
	// 读取服务器的响应
	for read {
		count, err = con.Read(data)
		read = (err == nil)
		fmt.Printf(string(data[0:count]))
	}
	con.Close()
}
