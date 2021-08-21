package main

import "fmt"

type Number struct {
	max []*MaxNum
	min int
}

type MaxNum struct {
	year  int
	month int
	day   int
}

//避免了空指针问题，如果为nil ，直接返回nil.
func (n *Number) GetMax() []*MaxNum {
	if n != nil {
		return n.max
	}
	return nil
}

func main() {
	var num *Number

	//maxs := num.max
	//for _, v := range maxs {
	//	fmt.Println(v.year)
	//}

	fmt.Println(num.GetMax())

}
