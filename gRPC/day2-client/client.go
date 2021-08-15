package day2_client

import (
	"fmt"
	"net"
)

//通用体
type call struct {
	reply string
	err   error
}

//客户端连接服务端
func dial() {
	//获取到连接对象
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Dial error")
	}

	//客户度发送数据
	conn.Write([]byte("abc"))

	var buf []byte
	conn.Read(buf)
}
