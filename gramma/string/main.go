package main

import (
	"fmt"
	"strings"
)

func main(){


	var styleList [] string
	styleList = append(styleList,"a","b","c")
	//根据sep将数组元素拼接成string，
	fmt.Println(strings.Join(styleList, "*"))

}