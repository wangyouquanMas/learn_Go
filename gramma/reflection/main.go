package main

import (
	"fmt"
	"reflect"
)

//反射是一种机制来检查 类型和value值存储在接口变量中的

//Package中 reflect type /value 可以获取到接口变量值的内容
//  reflect.TypeOf reflect.ValueOf

// TypeOf returns the reflection Type of the value in the interface{}.
//func TypeOf(i interface{}) Type
//When we call reflect.TypeOf(x), x is first stored in an empty interface,
//which is then passed as the argument; reflect.TypeOf unpacks that empty interface to recover the type information.

//方法功能
//One important example is that Value has a Type method that returns the Type of a reflect.Value.
//both Type and Value have a Kind method that returns a constant indicating what sort of item is stored: Uint, Float64, Slice,
func main() {
	var y float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(y))

	fmt.Println("value:", reflect.ValueOf(y).String())

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	//可以根据值获取到类型
	fmt.Println("type:", v.Type())
	//可以根据kind返回常量 来判断是什么类型
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	//根据根据value with names 获取值
	fmt.Println("value:", v.Float())

	var x1 uint8 = 'x'
	v1 := reflect.ValueOf(x1)
	fmt.Println("type:", v1.Type())
	fmt.Println("kind is uint8:", v1.Kind() == reflect.Uint8)
	x1 = uint8(v1.Uint())

	type MyInt int
	var x2 MyInt = 7
	v2 := reflect.ValueOf(x2)
	fmt.Println(v2.Type())
	fmt.Println(v2.Kind())

	/*
		2 reflect.value ： 使用接口方法恢复接口值
	*/
	//
	//y := v.Interface().(float64)

	fmt.Println(y)

	/*
	   3 reflect.New : 返回值的类型是指针类型
	*/

	//ReplyType reflect.Type

	//replyv := reflect.New(m.ReplyType.Elem())
	//switch m.ReplyType.Elem().Kind() {
	//case reflect.Map:
	//	replyv.Elem().Set()
	//}
}

type Method struct {
	ReplyType reflect.Type
}

type Value struct {
	a float64
	b float64
	c float64
}

func (v Value) Interface() interface{} {
	return v
}
