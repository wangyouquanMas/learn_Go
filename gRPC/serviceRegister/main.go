package serviceRegister

import (
	"fmt"
	"reflect"
)

type Server struct {
}

type ServiceDesc struct {
	ServiceName string
	//HandlerType: 用于检查接口是否被实现
	HandlerType interface{}
}

/*
	1 参数说明
		sd: 服务描述
        ss：服务实例
*/
func RegisterService(sd *ServiceDesc, ss interface{}) {
	ht := reflect.TypeOf(sd.HandlerType).Elem()
	fmt.Println()
}
