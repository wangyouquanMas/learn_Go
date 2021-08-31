package main

import (
	"fmt"
	"sort"
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

	//字符串切换，不区分中英文
	s2 := "abc"
	res := "["
	for i := 0; i < len([]rune(s2)); i++ {
		//ch := []rune(s2)[i]
		prefix := "<em class=\"keyword\">"
		suffix := "</em>"
		temp := prefix + string([]rune(s2)[i]) + suffix
		res = res + temp
	}
	res = res + "]"
	fmt.Println(res)

	//if req.NameRaw != "" {
	//	res := "["
	//	for i := 0; i < len(up.Name); i++ {
	//		prefix := "<em class=\"keyword\">"
	//		suffix := "</em>"
	//		temp := prefix + string([]rune(up.Name)[i]) + suffix
	//		res = res+temp
	//		//up.Name = "[<em class=\"keyword\">l</em><em class=\"keyword\">e</em><em class=\"keyword\">a</em><em class=\"keyword\">s</em><em class=\"keyword\">o</em><em class=\"keyword\">n</em>]"
	//	}
	//	res = res+"]"
	//}

	//strings.LastIndex 返回子串在 str中最后一次出现的索引位置
	//serviceMethod := "Search.GetUpRes"
	//dot := strings.LastIndex(serviceMethod, ".")
	//fmt.Println(dot)

	//fmt.Println(serviceMethod[:dot] + "," + serviceMethod[dot+1:])
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
	//var s = "abc"
	//fmt.Println(strings.Fields(s))
	//fmt.Println([]string{s})

	// 字符串排序
	//
	//s2 := []string{"eat", "tea", "ate"}
	//
	//sort.Strings(s2)
	//
	//fmt.Println(s2)

	w1 := "bcad"
	w2 := SortString(w1)

	fmt.Println(w1)
	fmt.Println(w2)
}

//}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

//func main() {
//
