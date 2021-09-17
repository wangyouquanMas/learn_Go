package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	server := ":7373"
	netListen, err := net.Listen("tcp", server)
	if err != nil {
		fmt.Println("connect err:", err)
		os.Exit(1)
	}
	fmt.Println("waiting for client.....")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), "fatal error:", err)
			continue
		}
		//设置短连接 (10s)
		conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))

		fmt.Println(conn.RemoteAddr().String(), "connect success!")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), "fatal error:", err)
			return
		}

		Data := buffer[:n]
		message := make(chan byte) //?数据从哪里写入？
		//心跳计时
		go HeartBeating(conn, message, 4)
		//检测每次是否有数据传入
		go GravelChannel(Data, message)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.0000000"), conn.RemoteAddr().String(), string(buffer[:n]))
	}
}

//每次连接的携程拥有的message 都是独立的
func GravelChannel(bytes []byte, mess chan byte) {
	for _, v := range bytes {
		mess <- v
	}
	close(mess)
}

func HeartBeating(conn net.Conn, bytes chan byte, timeout int) {
	select {
	case fk := <-bytes:
		fmt.Println(conn.RemoteAddr().String(), "心跳：第", string(fk), "times")
		conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break

	case <-time.After(5 * time.Second):
		fmt.Println("conn dead now")
		conn.Close()
	}
}
