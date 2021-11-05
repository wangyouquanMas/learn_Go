package main

import "fmt"

/*
 new
   返回值为指向该类型内存地址的指针
    同时会把分配的内存置为零，也就是类型的零值, 即字符为空，整型为0, 逻辑值为false
*/

type test struct {
	age  int
	name string
}

func main() {

	t := new(test)

	fmt.Println(t, " ", *t)

	//make 仅用来分配及初始化类型为 slice、map、chan 的数据  ; slice 0值时nil, make之后初始化slice 中元素值为默认值
	m := make([]int64, 5, 5)
	fmt.Println(m)
}
