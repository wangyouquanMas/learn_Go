package main

import (
	"encoding/json"
	"fmt"
)

/*
1 :Json Marshal：将数据编码成json字符串
*/
type Stu struct {
	Name string `json:"name"`
	Age  int
}

func main() {

	stu := &Stu{
		Name: "Apple",
		Age:  20,
	}
	stuByte, _ := json.Marshal(stu)
	fmt.Println(string(stuByte))

	jsonInpu = ""

	json.Unmarshal(stuByteUm, *stu)

	fmt.Println(string(stuByteUm))
}
