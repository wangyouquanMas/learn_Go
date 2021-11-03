package main

import (
	"log"
	"net"
	"os"
	"time"
)

/*
   conn.SetReadDeadline(time.Now().Add(readTimeout))
	设置conn.read超时时间

  超时场景：1 等待服务端处理时，等待处理导致的超时
      server 写入conn 耗时 3s , client read超时等待2s，就会导致read timeout
*/

func main() {
	connTimeout := 10 * time.Second
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8080", connTimeout) // 3s timeout
	if err != nil {
		log.Println("dial failed:", err)
		os.Exit(1)
	}
	defer conn.Close()

	readTimeout := 10 * time.Second

	buffer := make([]byte, 512)

	for {
		err = conn.SetReadDeadline(time.Now().Add(readTimeout)) // timeout
		if err != nil {
			log.Println("setReadDeadline failed:", err)
		}

		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Read failed:", err)
			//break
		}
		log.Println("count:", n, "msg:", string(buffer))
	}

}
