package main

import "fmt"

type StructA struct {
	*StructB
}

type StructB struct {
}

func (b StructB) testb() {
	fmt.Println("struct a 直接调用 内嵌的struct b 的方法")
}

func main() {
	b := &StructB{}
	a := &StructA{b}
	a.testb()
}
