package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
	io.copy作用：在文件指针之间直接复制

	适用场景：下载大文件或者复制大文件。 因为io.copy就是在文件指针之间直接复制的,不用将文件读入内存
*/

func main() {
	url := "http://baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, "get url error", err)
	}
	defer resp.Body.Close()

	out, err := os.Create("/tmp/baidu.text")
	wt := bufio.NewWriter(out)
	defer out.Close()
	n, err := io.Copy(wt, resp.Body)
	fmt.Println("write", n)
	if err != nil {
		panic(err)
	}
	wt.Flush()
}
