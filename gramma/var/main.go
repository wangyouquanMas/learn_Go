package main

import "fmt"

//变量定义

func main() {
	//1定义变量并初始化值
	//初始化“variableName”的变量为“value”值，类型是“type”
	//var variableName type = value

	//2同时初始化多个变量
	/*
	   定义三个类型都是"type"的变量,并且分别初始化为相应的值
	   vname1为v1，vname2为v2，vname3为v3
	*/
	//var vname1, vname2, vname3 type= v1, v2, v3

	//3以直接忽略类型声明
	/*
	   定义三个变量，它们分别初始化为相应的值
	   vname1为v1，vname2为v2，vname3为v3
	   然后Go会根据其相应值的类型来帮你初始化它们
	*/
	//var vname1, vname2, vname3 = v1, v2, v3

	//4 简短声明 用在函数内部
	//编译器会根据初始化的值自动推导出相应的类型

	//vname1, vname2, vname3 := v1, v2, v3

	//5 _（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。在这个例子中，我们将值35赋予b，并同时丢弃34：
	//_, b := 34, 35

	//6 已声明但未使用的变量会在编译阶段报错

	res := test()
	fmt.Println(res)
}

type Item struct {
	a int
}

func test() (res *Item) {

	res, err := test1()
	if err != nil {
		return
	}
	fmt.Println(res)
	return
}

func test1() (res *Item, err error) {
	return &Item{
		a: 3,
	}, nil
}
