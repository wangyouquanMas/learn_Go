package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

/*
   类型断言
      功能：
           1 检查 i 是否为 nil
           2 检查 i 存储的值是否为某个类型
*/

//  1 检查 i 是否为 nil
//t := i.(T)
//这个表达式可以断言一个接口对象（i）里不是 nil，并且接口对象（i）存储的值的类型是 T，如果断言成功，就会返回值给 t，
// 如果断言失败，就会触发 panic。
//func main() {
//	//接口对象（i)
//	var i interface{} //i为nil时，断言失败
//	//var i interface{} = 10
//	//断言成功，就会返回值给 t1
//	t1 := i.(int)
//	fmt.Println(t1)
//
//	fmt.Println("=====分隔线=====")
//
//	//断言失败，触发 panic。
//	t2 := i.(string)
//	fmt.Println(t2)
//
//}

//检查 i 存储的值是否为某个类型
//t, ok:= i.(T)
//如果断言失败，但和第一种表达式不同，这个不会触发 panic，而是将 ok 的值设为 false ，表示断言失败，此时t 为 T 的零值。
//表达式也是可以断言一个接口对象（i）里不是 nil，并且接口对象（i）存储的值的类型是 T，如果断言成功，就会返回其类型给 t

func main() {
	var i interface{} = 10
	t1, ok := i.(int)
	fmt.Printf("%d-%t\n", t1, ok)

	fmt.Println("=====分隔线1=====")

	t2, ok := i.(string) // 断言失败，t2为 string 的零值
	fmt.Printf("%s-%t\n", t2, ok)

	fmt.Println("=====分隔线2=====")

	var k interface{} // nil
	t3, ok := k.(interface{})
	fmt.Println(t3, "-", ok)

	fmt.Println("=====分隔线3=====")
	k = 10
	t4, ok := k.(interface{})
	fmt.Printf("%d-%t\n", t4, ok)

	t5, ok := k.(int)
	fmt.Printf("%d-%t\n", t5, ok)

	fmt.Println("=====分隔线4=====")
	// interface也是一个类型， interface不是任意类型
	type Tp struct {
		Strs interface{} `json:"strs"`
	}
	var k1 = &Tp{}
	var k8 interface{} = []string{"1", "3"}
	fmt.Printf("k8类型为 （%+v)", reflect.TypeOf(k8))
	fmt.Println()
	//var k8 interface{}
	//k8 = []string{"1", "3"}
	//fmt.Println(k8.([]string)[0])

	fmt.Printf("unmarshal前 k1.Strs类型为 （%+v)", reflect.TypeOf(k1.Strs))
	_ = json.Unmarshal([]byte(`{"strs":["1","2"]}`), &k1)
	fmt.Println()
	fmt.Printf("unmarshal后 k1.Strs类型为 （%+v)", reflect.TypeOf(k1.Strs))
	fmt.Println()
	if t6, ok := k1.Strs.([]interface{}); ok {
		s := t6[0].(string)
		fmt.Println(s)
	} else {
		fmt.Println("false")
	}

	fmt.Println("=====分隔线5=====")
	a := reflect.TypeOf("")
	fmt.Println(a.Kind())
	argv := reflect.New(reflect.TypeOf(""))
	// reflect.New： 返回指针
	fmt.Println(argv.Kind())
	//argv.Interface(): 转为接口类型 类似  var i interface = argv
	fmt.Println(argv.Interface())
}

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

var w io.Writer

func test(r io.Writer) {
	value, ok := r.(io.Writer)
	if !ok {
		fmt.Println("It's not ok for type io.Writer")
		return
	}
	var p []byte
	_, err := value.Write(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The value is ", string(p))
}

//
//func main() {
//	tty, _ := os.OpenFile("/Users/wang/example.go", os.O_RDWR, 0)
//	// r包含(value, value type) pair ，（tty,*os.File） ; 当然 os.File不止实现了 read方法 。
//	// 尽管 value 仅提供了获取read method权限 。 但是在value值中携带着关于该值的所有类型信息 。
//
//	test(tty)
//
//	// 可以传入tty类型 *os.File ， 也可以传入tty实现的接口之一【这里就是io.writer】
//	//value,ok:=r.(*os.File) 也可以。
//	// 可以获取到 tty所有支持的方法。
//	r = tty
//	// r接口仅提供了 tty (os.File)实现的 writer方法
//	// 该assert 表明 tty实现了io.writer 接口。 也可以直接赋值给 io.Writer对象
//	//var w io.Writer
//	//w = r.(io.Writer)
//	//w 包含了 （tty, *os.File）和 r中拥有相同的 pair
//	//The static type [定义在 接口中的方法 ]of the interface determines what methods may be invoked with an interface variable,
//	//even though the concrete value 【tty】inside may have a larger set of methods [*os.File类型 实现了其他方法]
//
//	// 空接口  包含相同的 pair(tty,*os.File)
//	// 空接口可以持有任何值 以及该值的所有信息
//	//var empty interface{}
//	//empty = w
//
//	//该assert 表明 tty实现了io.writer 接口。
//	value, ok := r.(io.Writer)
//	if !ok {
//		fmt.Println("It's not ok for type io.Writer")
//		return
//	}
//	content := "Goodbye"
//	_, err := value.Write([]byte(content))
//	if err != nil {
//		fmt.Println(err)
//	}
//	//fmt.Println("The value is ", )
//	//fmt.Println("The value is ", )
//
//	tty = r.(*os.File)
//
//	/*1 reflect.TypeOf():可以接收的任意类型*/
//	//数值
//	var i = 3
//	intType := reflect.TypeOf(i)
//	//intType.Kind():返回元素的基本类型
//	fmt.Println(intType.Kind())
//	//intType.NumMethod()：计算当前类型包含方法数
//	fmt.Println(intType.NumMethod())
//	//
//	fmt.Println(intType.Name())
//	//返回分配内存的字节数
//	fmt.Println(intType.Align())
//	fmt.Println(intType.PkgPath())
//	//返回函数的输入参数
//	//fmt.Println(intType.NumIn())
//
//	fmt.Println("*********指针************")
//	//指针
//	var t = 3
//	var i1 *int
//	i1 = new(int)
//	i1 = &t
//	i1P := reflect.TypeOf(i1)
//	fmt.Println(i1P.Elem())
//	fmt.Println(i1P.NumMethod())
//	fmt.Println(i1P.Name())
//	fmt.Println(i1P.Align())
//
//	fmt.Println("*********数组************")
//	//数组
//	var arr = []int{1, 2, 3}
//	reflect.TypeOf(arr)
//	fmt.Println(intType.NumMethod())
//
//	fmt.Println("*********切片************")
//	//切片
//	var slice []int
//	slice = make([]int, 3)
//	reflect.TypeOf(slice)
//	fmt.Println(intType.NumMethod())
//
//	fmt.Println("*********map************")
//	//map
//	var mapT map[int]string
//	mapT = make(map[int]string, 3)
//	reflect.TypeOf(mapT)
//	fmt.Println(intType.NumMethod())
//
//	fmt.Println("*********chan************")
//	//chan
//	//var chanT chan int
//	//chanT = make(chan int, 3)
//	//chanT= reflect.TypeOf(chanT)
//	fmt.Println(intType.NumMethod())
//
//	fmt.Println("*********结构体************")
//	//struct
//	//structT := structT{
//	//	a :1,
//	//	b:1,
//	//}
//	structT := new(structT)
//	structA := reflect.TypeOf(structT)
//	//NumMethod： 对于非接口类型，只计算可导出方法数量；
//	fmt.Println(structA.NumMethod())
//	//fmt.Println(structA.Method(0).Name)
//	//fmt.Println(structA.NumMethod())
//
//	fmt.Println("*********函数************")
//	//函数
//	funcA := func(args string) {
//		fmt.Println(args)
//	}
//	funcA("hello world")
//	funcT := reflect.TypeOf(funcA)
//	fmt.Println(funcT.NumIn())
//	fmt.Println(funcT.NumOut())
//	fmt.Println(funcT.NumMethod())
//
//	fmt.Println("*********接口************")
//	var interA interfaceA
//	var interB interfaceB
//	//指针初始化
//	interB = structT
//	interA = interB
//	interfaceT := reflect.TypeOf(interA)
//	fmt.Println(interfaceT.NumMethod())
//
//	fmt.Println("*********nil************")
//	//reflect.TypeOf((*error)(nil)).Elem(): 返回类型元素的名称nil
//	fmt.Println(reflect.TypeOf((*error)(nil)).Elem())
//}
//
//type structT struct {
//	a int
//	b int
//}
//
////func (s structT) Test1(str string) (out string) {
////	return str
////}
//
//type interfaceB interface {
//	test1()
//	test2()
//	test3()
//}
//
//type interfaceA interface {
//	test1()
//	test2()
//}
//
//func (s structT) test1() {
//
//}
//
//func (s structT) test2() {
//
//}
//
//func (s structT) test3() {
//
//}

//TODO : 断言的使用
//func (s *Service) getUpSearchHighlight(hl *model.UpperElastic, up *pb.ArchiveLangUp) (err error) {
//	if hl == nil || up == nil {
//		return
//	}
//	if names, exist := hl.Name.([]interface{}); exist {
//		if len(names) > 0 {
//			up.Name = names[0].(string)
//		}
//	}
//	return
//}
