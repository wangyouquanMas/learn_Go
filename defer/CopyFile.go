package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close() //保障文件最终关闭

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	defer dst.Close()

	//written, err = io.Copy(dst, src)
	//dst.Close()
	//src.Close()
	return  io.Copy(dst, src)
}

func a() (res int) {
	i := 0
	defer fmt.Println(i)
	i++
	return i
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func write(fileName ,text string) error{

	file,err:=os.Create(fileName)
	if err != nil{
		return err
	}
	defer file.Close()

	_,err = io.WriteString(file,text)
	if err != nil{
		return err
	}

	return file.Close()

}

func fileCopy(source,destination string) error{

	src,err:= os.Open(source)
	if err!=nil{
		return err
	}

	defer src.Close()

	dst,err:=os.Create(destination)
	if err!=nil{
		return err
	}

	defer dst.Close()

	n,err:=io.Copy(dst,src)
	if err!=nil{
		return err
	}

	fmt.Printf("Copied %d bytes frm  %s to %s\n",n,source,destination)

	return dst.Close()


}

func main(){

	res,_:=CopyFile("/Users/wang/b.txt","/Users/wang/a.txt")
    fmt.Println(res)


	i :=a()
	fmt.Println(i)

	b()

	c()


	if err := write("sample.txt", "This file contains some sample text."); err != nil {
		log.Fatal("failed to create file")
	}

	if err := fileCopy("sample.txt", "sample-copy.txt"); err != nil {
		log.Fatal("failed to copy file: %s")
	}
}


