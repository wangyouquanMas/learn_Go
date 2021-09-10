package main

import "fmt"

type Number struct {
	max *MaxNum
	min int
}

type MaxNum struct {
	year  int
	month int
	day   int
}

//避免了空指针问题，如果为nil ，直接返回nil.
//func (n *Number) GetMax() []*MaxNum {
//	if n != nil {
//		return n.max
//	}
//	return nil
//}

func main() {
	//var num *Number
	//
	////maxs := num.max
	////for _, v := range maxs {
	////	fmt.Println(v.year)
	////}
	//
	//fmt.Println(num.GetMax())

	//指针必须要先初始化，没有初始化就返回，会报错 ？？
	//fmt.Println(Pointer(2))

	//已经初始化指针类型
	var test *Number
	test = &Number{}
	test.max = &MaxNum{}

	//test = testInit()
	fmt.Println(test.max)
	fmt.Println(test.min)

	//申明指针变量，没有初始化，直接复制
	var a *int
	b := 7
	a = &b

	fmt.Println(a)

	//1 初始化为null,直接赋初值
	Number := &Number{}
	Number = testInit()
	fmt.Println(Number)
}

func testInit() (num *Number) {

	num = &Number{}

	num = &Number{
		min: 1,
	}

	return num

}

func Pointer(a int) (res *MaxNum, e error) {

	if a < 3 {
		return
	}

	res = &MaxNum{
		year: 2021,
	}

	return res, e
}
