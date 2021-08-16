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
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))

	fmt.Println("value:", reflect.ValueOf(x).String())

}
