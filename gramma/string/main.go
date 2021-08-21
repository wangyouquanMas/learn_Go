package main

import (
	"fmt"
	"strings"
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
	////fmt.Print(len(word))

	//var a = "abc"

	//res, _ := strconv.Atoi(a)
	//fmt.Println(res)

	//for i, v := range a {
	//	fmt.Println(i, v)
	//}

	// 字符串遍历

	//for i := 0; i < len(a); i++ {
	//	ch := a[i]
	//	//ctype := reflect.TypeOf(ch)
	//	//fmt.Printf("%s", ctype)
	//	fmt.Printf(string(ch))
	//}

	// 字符串拼接
	//var res []string
	//res = make([]string, 0)
	//
	//var s []string
	//s = make([]string, 0)
	//s = append(s, "abc")
	//var s1 string
	//s1 = "def"
	//
	//k := 0
	//for i := 0; i < len(s1); i++ {
	//	for j := 0; j < len(s[k]); j++ {
	//		res = append(res, strings.Join([]string{string(s1[i]), string(s[k][j])}, ""))
	//	}
	//}
	//
	//fmt.Println(res)

	//
	var s = "abc"
	fmt.Println(strings.Fields(s))
	fmt.Println([]string{s})
}
