package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a = []string{"a", "b", "c"}
	b, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	var result = string(b)
	fmt.Println(result)
}
