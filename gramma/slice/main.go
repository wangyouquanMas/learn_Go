package main

import "fmt"

// slice

//在初始定义数组时，我们并不知道需要多大的数组，因此我们就需要“动态数组”。在Go里面这种数据结构叫slice
func main() {

	//1 slice并不是真正意义上的动态数组，而是一个引用类型 [会改变值]。slice总是指向一个底层array
	//slice的声明也可以像array一样，只是不需要长度。
	// 和声明array一样，只是少了长度
	//var fslice []int

	//2声明一个slice，并初始化数据
	//slice := []byte{'a', 'b', 'c', 'd'}

	//3 slice可以从一个数组或一个已经存在的slice中再次声明。
	//slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i。
	// 声明一个含有10个元素元素类型为byte的数组
	//var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	//var a, b []byte
	//a = ar[2:5]

	//4 slice的默认开始位置是0，ar[:n]等价于ar[0:n]

	//5 slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]

	//6 如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]

	//7 对于slice有几个有用的内置函数
	/*	len 获取slice的长度
		cap 获取slice的最大容量
		append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
		copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数*/

	//数组指针传递
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(sum(&arr))
	fmt.Println(arr)
}

func sum(arr *[5]int) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	return s
}
