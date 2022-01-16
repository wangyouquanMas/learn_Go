package main

import (
	"encoding/json"
	"fmt"
)

/*
1 :Json Marshal：将数据编码成json字符串
2 :  marshol的结构体 映射是根据字段名称来进行匹配的

   比如在stu1中 对 Name, Age进行marshal, 然后unmarshal stu2时， Name,Age字段会被映射，而Test没有，所以为默认值。
   也即是说， 支持字段个数不同的结构体的unmarshal， 只要有同名字段[区分大小写]，就会被映射
*/

func main() {

	type Stu1 struct {
		Name string `json:"name"`
		Age  int
	}

	type Stu2 struct {
		Name string `json:"name"`
		Age  int
		Test int
	}

	stu3 := &Stu1{
		Name: "Apple",
		Age:  20,
	}

	var stu41 Stu2

	stuByte, _ := json.Marshal(stu3)
	fmt.Println(string(stuByte))

	json.Unmarshal(stuByte, &stu41)
	//json.Unmarshal(stuByte, *stu41)
	fmt.Println((stu41))
}
