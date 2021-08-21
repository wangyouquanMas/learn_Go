package main

import (
	"fmt"
	"io"
	"os"
)

//func main() {
//	reflection()
//}
//
//func reflection() (bool, error) {
//	var r io.Reader
//	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
//	if err != nil {
//		return nil, err
//	}
//	r = tty
//}

//类型和接口

type MyInt int

// 尽管i,j 都有潜在的类型，它们之间也必须要进行强制转换才能赋值
var i int
var j MyInt

var r io.Reader

//var w io.Writer

//func test(r io.Writer) {
//	value, ok := r.(io.Writer)
//	if !ok {
//		fmt.Println("It's not ok for type io.Writer")
//		return
//	}
//	var p []byte
//	_, err := value.Write(p)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("The value is ", string(p))
//}

func main() {
	tty, _ := os.OpenFile("/Users/wang/example.go", os.O_RDWR, 0)
	// r包含(value, value type) pair ，（tty,*os.File） ; 当然 os.File不止实现了 read方法 。
	// 尽管 value 仅提供了获取read method权限 。 但是在value值中携带着关于该值的所有类型信息 。

	//test(tty)

	// 可以传入tty类型 *os.File ， 也可以传入tty实现的接口之一【这里就是io.writer】
	//value,ok:=r.(*os.File) 也可以。
	// 可以获取到 tty所有支持的方法。
	r = tty
	// r接口仅提供了 tty (os.File)实现的 writer方法
	// 该assert 表明 tty实现了io.writer 接口。 也可以直接赋值给 io.Writer对象
	var w io.Writer
	w = r.(io.Writer)
	//w 包含了 （tty, *os.File）和 r中拥有相同的 pair
	//The static type [定义在 接口中的方法 ]of the interface determines what methods may be invoked with an interface variable,
	//even though the concrete value 【tty】inside may have a larger set of methods [*os.File类型 实现了其他方法]

	// 空接口  包含相同的 pair(tty,*os.File)
	// 空接口可以持有任何值 以及该值的所有信息
	var empty interface{}
	empty = w

	//该assert 表明 tty实现了io.writer 接口。
	value, ok := r.(io.Writer)
	if !ok {
		fmt.Println("It's not ok for type io.Writer")
		return
	}
	content := "Goodbye"
	_, err := value.Write([]byte(content))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("The value is ", )
	//fmt.Println("The value is ", )

	tty = r.(*os.File)

}
