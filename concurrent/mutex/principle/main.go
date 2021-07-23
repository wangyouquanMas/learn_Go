package main

import "sync"

//Mutex锁的实现过程

//1 初步：flag实现Mutex
//  特点： flag字段标记goroutine是否持有锁,如果为true就持有该锁，其他goroutine就等待，如果为false，就通过cas将该值设置为true。
//  存在问题，请求锁的goroutine会排队等待获取互斥锁，从性能角度来看，如果将锁交给正在占用CPU时间片的goroutine，就不用做上下文切换，高并发场景下会有更好的性能
var flag bool

var mu sync.Mutex

func main() {

	for {

		if flag == true {
			break
		}
	}

	go resources()
	go resources()

}

func resources() {
	flag = true
	var i int
	i++
	flag = false
}
