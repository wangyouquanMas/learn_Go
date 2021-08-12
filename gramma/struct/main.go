package main

import (
	"fmt"
	"log"
	"net/http"
)

//结构体的实例化，会让其作为指针接收器的方法也实例化？

type Dog struct {
	name string
}

//结构体指针实现Handler接口，所以 new(Dog)是实现了接口的实例化对象
func (d *Dog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.name = "dog"
	fmt.Printf("%s barking", d.name)
}

//1结构体内元素定义
type Man struct {
	name    string
	age     int
	company Company
}

type Company struct {
	name     string
	industry int
}

func main() {

	//1 结构体初始化方式
	c := Man{
		name: "wang",
		age:  24,
		company: Company{
			name:     "wangdeguo",
			industry: 2,
		},
	}

	fmt.Println(c)

	//type Handler interface {
	//	ServeHTTP(ResponseWriter, *Request)
	//}
	//在其他的函数中，将该interface类型作为函数的形参： func Handle(pattern string, handler Handler)
	http.Handle("/count", new(Dog))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
