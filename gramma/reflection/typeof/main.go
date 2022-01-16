package main

import (
	"fmt"
	"reflect"
)

/*
	TypeOf returns the reflection Type that represents the dynamic type of i.

     那么这个动态类型包含了什么？

     以下面实例为例：Arith作为方法接受者， 通过对该类型进行实例化后，使用该实例就可以调用该方法。
	而通过reflect.typeof 返回的就是 该类型 *main.Arith ，通过该类型就可以获取到方法等其它属性，同时通过方法处理，也可以获取到一些其他信息。。。
*/

type Arith int

func (a *Arith) Add() {
	fmt.Println("1+1=2")
}

func main() {
	arith := new(Arith)
	fmt.Println(reflect.TypeOf(arith))
}
