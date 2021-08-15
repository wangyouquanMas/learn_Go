package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	fmt.Println("tcp server run...")
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer listener.Close()
	var num int
	for {
		//永久循环来接受来自客户端的连接，accept()会等待并返回一个客户端的连接:
		// 堵塞
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Println("accept err:", err2)
		}
		fmt.Println(num)
		num++
		//每个连接都必须创建新线程（或进程）来处理，否则，单线程在处理连接的过程中，无法接受其他客户端的连接
		go handleConn(conn)
	}
}

//网络流都要关闭
func handleConn(conn net.Conn) {
	defer conn.Close()
	remoteAddr := conn.RemoteAddr()
	fmt.Println(remoteAddr, "connect success")
	//接收数据
	//预先申请[]byte类型的变量 ，用于存放 从conn对象读取的数据。
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		if "exit" == string(buf[:n-2]) {
			fmt.Println(remoteAddr, "已断开")
			return
		}
		fmt.Printf("from %s data:%s\n", remoteAddr, string(buf[:n]))
		to := strings.ToUpper(string(buf[:n]))

		//通过Write()方法再写入到conn对象中。
		conn.Write([]byte(to))
	}
}
