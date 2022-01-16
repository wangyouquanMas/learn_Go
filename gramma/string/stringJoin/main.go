package main

import "fmt"

func main() {
	//str := []string{"{2021122916777216}", "" + strconv.Itoa(1), "" + strconv.Itoa(2)}
	//res := strings.Join(str, "-")
	//fmt.Println(res)

	str := "520"
	fmt.Println([]byte(str))
	by := []byte{231, 136, 177, 230, 131, 133}

	fmt.Println(string(by))
}
