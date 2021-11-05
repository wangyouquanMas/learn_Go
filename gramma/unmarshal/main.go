package main

import (
	"encoding/json"
	"fmt"
)

/*
1 :Json Marshal：将数据编码成json字符串
*/
type Stu struct {
	Name  string `json:"name"`
	Age   int
	style int
}

type Stu1 struct {
	Name  string `json:"name"`
	Age   int
	style int
}

type Stu2 struct {
	Name string `json:"name"`
	Age  int
}

func main() {

	stu := &Stu1{
		Name: "Apple",
		Age:  20,
	}

	//stu2 := &Stu{
	//	Name:  "Apple",
	//	Age:   20,
	//	style: 1,
	//}

	stuByte, _ := json.Marshal(stu)
	fmt.Println(string(stuByte))

	json.Unmarshal(stuByte, *stu)

	fmt.Println((stu))
}
