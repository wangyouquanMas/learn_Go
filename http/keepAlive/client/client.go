package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	server := "127.0.0.1:7373"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Println("fatal error", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("fatal error", err.Error())
		os.Exit(1)
	}

	fmt.Println(conn.RemoteAddr().String(), "connection success!")

	sender(conn)
	fmt.Println("sender over")
}

func sender(conn *net.TCPConn) {
	for i := 0; i < 10; i++ {
		words := strconv.Itoa(i) + "Hello I'm MyHeartBeat Client."
		fmt.Println([]byte(words))
		msg, err := conn.Write([]byte(words))
		if err != nil {
			fmt.Println("fatal error:", err)
			os.Exit(1)
		}
		fmt.Println("服务端接收了", msg)
		time.Sleep(2 * time.Second)
	}

	for i := 0; i < 2; i++ {
		time.Sleep(12 * time.Second)
	}

	for i := 0; i < 10; i++ {
		words := strconv.Itoa(i) + "Hello I'm MyHeartBeat Client."
		fmt.Println([]byte(words))
		msg, err := conn.Write([]byte(words))
		if err != nil {
			fmt.Println("fatal error:", err)
			os.Exit(1)
		}
		fmt.Println("服务端接收了", msg)
		time.Sleep(2 * time.Second)
	}
}
