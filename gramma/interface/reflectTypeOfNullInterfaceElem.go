package main

import (
	"fmt"
	"reflect"
)

/*
	1 reflect.TypeOf(i interface{}) 功能
		返回i的动态类型
		如果输入是空接口，那么会返回nil

	注意：关于空接口的说明
		一个接口包含了两个属性 type,value，只有type,value都为nil，接口才为nil

	2 reflect.TypeOf(nullInterface).Elem() 功能
		返回nullInterface类型的基本类型，如*main.SearchServer 返回main.SearchServer
		*error  返回error
*/

type SearchServer interface {
}

func main() {
	var nullInterface interface{}
	nullInterface = (*SearchServer)(nil)
	nullInterfaceType := reflect.TypeOf(nullInterface).Elem()
	nullInterfaceValue := reflect.ValueOf(nullInterface)
	fmt.Println(nullInterfaceType, nullInterfaceValue)
}
