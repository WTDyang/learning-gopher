package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("udp", "127.0.0.1:52148")
	defer func() {
		conn.Close()
		fmt.Println("客户端已退出")
	}()

	//客户端发起交谈
	conn.Write([]byte("hello server, i am client"))

	////接收服务端消息
	//buffer := make([]byte, 1024)
	//n, _ := conn.Read(buffer)
	//
	//fmt.Println("服务端：" + string(buffer[:n]))

}
