package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	client, _ := net.Dial("tcp", "127.0.0.1:8080")

	//当前的go routine 堵塞了 ,只有server解析完请求后，才退出执行主协程。

	go func() {
		input := make([]byte, 1024)

		for {
			n, err := os.Stdin.Read(input)
			if err != nil {
				fmt.Println("input err:", err)
				continue
			}
			client.Write([]byte(input[:n]))
			fmt.Println("第一个client : ", client)
		}
	}()

	buf1 := make([]byte, 1024*1024)
	n, _ := client.Read(buf1)
	fmt.Println("第二个client :", string(buf1[:n]))

	//所以 client.read 不会写入客户端输入的数据client.Read(buf1)。
	buf := make([]byte, 1024)
	for {
		n, err := client.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("read err:", err)
			continue
		}
		fmt.Println(string(buf[:n]))
	}
}
