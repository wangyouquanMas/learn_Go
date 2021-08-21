package main

import (
	"encoding/json"
	"fmt"
)

/*
   Unmarshal 解析 json encoded 的byte，赋值给 v
v 不可以为nil ，要为指针 否为会报异常
   v 就是原先的被 编码为 byte的值
*/

func main() {

	var nameIdMap map[int]int
	var err error
	var mJson []byte
	var umJson map[int]int
	nameIdMap = make(map[int]int, 0)

	for i := 0; i < 10; i++ {
		nameIdMap[i] = i
	}

	if mJson, err = json.Marshal(nameIdMap); err == nil {
		fmt.Println(string(mJson))
	}

	if err = json.Unmarshal(mJson, &umJson); err == nil {
		fmt.Println(umJson)
	}

	for i, v := range umJson {
		fmt.Println(i, v)
	}
}
