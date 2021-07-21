package main

import "fmt"

//1 method "继承"
// method也是可以 "继承" 的。如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。

type A struct {
	name  string
	age   int
	phone string
}

type B struct {
	A      //匿名字段
	school string
}

type C struct {
	A       //匿名字段
	company string
}

func (a A) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", a.name, a.phone)
}

//Employee的method重写Human的method ，重写了匿名字段的方法
func (e *C) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
	mark := B{A{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := C{A{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}
