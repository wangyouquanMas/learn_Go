package main

/*
	new关键字作用：这个参数是一个类型，分配好内存后，返回一个指向该类型内存地址的指针[也就是要用指针去接收]。它同时把分配的内存置为零，也就是类型的零值。
			func new(Type) *Type
		对于引用类型的变量，不光要声明还要分配内容空间；

    make :用于内置引用类型（切片、map 和管道）的内存分配和初始化成员值，返回对象，而非指针。
*/

import (
	"fmt"
	"sync"
)

//func main() {
//	var i *int
//	i = new(int)
//	*i = 10
//	fmt.Println(*i)
//
//}

// new同时把分配的内存置为零，也就是类型的零值的好处 ：
//   user 类型中的lock字段不用初始化，直接就可以使用，因为已经在new时，被默认为零值
func main() {
	u := new(user)
	u.lock.Lock()
	u.name = "张三"
	u.lock.Unlock()

	fmt.Println(u)

	//make: 用于内置引用类型（切片、map 和管道）的内存分配和初始化成员值，返回对象，而非指针。
	arr := make([]*user, 3)
	mapT := make(map[int]string, 3)
	chanT := make(chan int, 3)

	//mew: 输入类型，为其分配零值内存，返回类型指针【也就是要用指针去接收】
	// new 用于值类型和用户定义的类型
	var intT *int
	intT = new(int)
	*intT = 10
	fmt.Println(*intT)
}

type user struct {
	lock sync.Mutex
	name string
	age  int
}
