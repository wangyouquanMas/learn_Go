package main

import (
	"fmt"
)

func main() {

	//var styleList []string
	//styleList = append(styleList, "a", "b", "c")
	////根据sep将数组元素拼接成string，
	//fmt.Println(strings.Join(styleList, "*"))
	//
	//// len
	//
	//word := "/lallala/b/c"
	//
	//fmt.Printf("%c", word[len(word)-1])
	//fmt.Print(len(word))

	//string -> []byte
	a := "{2021082416777216}-2-99"
	fmt.Println([]byte(a))

	//字符串比较
	s := "abc"
	var s1 = []string{"abc"}

	fmt.Println(s1[0] == s)
}
