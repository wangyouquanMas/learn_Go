package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(letterCombinations("234"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	if digits == "" || len(digits) == 0 {
		return nil
	}

	var numCharMap map[int]string
	numCharMap = make(map[int]string)
	numCharMap[2] = "abc"
	numCharMap[3] = "def"
	numCharMap[4] = "ghi"
	numCharMap[5] = "jkl"
	numCharMap[6] = "mno"
	numCharMap[7] = "pqrs"
	numCharMap[8] = "tuv"
	numCharMap[9] = "wxyz"

	if len(digits) == 1 {
		k, _ := strconv.Atoi(digits)
		s := numCharMap[k]
		var res []string
		for i := 0; i < len(s); i++ {
			res = append(res, string(s[i]))
		}
		return res
	}

	var s []string
	s = make([]string, 0)

	num, _ := strconv.Atoi(digits)
	var res1 []int
	for num != 0 {
		res1 = append(res1, num%10)
		num /= 10
	}
	for _, v := range res1 {
		//twoCombine(numCharMap[v], str)
		s1 := numCharMap[v]
		twoCombine(s1, &s)
	}
	return s
}

func twoCombine(s1 string, s *[]string) {
	if len(*s) == 0 {
		*s = append(*s, s1)
		return
	} else {
		var res []string
		res = make([]string, 0)
		if len(*s) == 1 {
			for i := 0; i < len(s1); i++ {
				for k := 0; k < len(*s); k++ {
					for j := 0; j < len((*s)[k]); j++ {
						res = append(res, strings.Join([]string{string(s1[i]), string((*s)[k][j])}, ""))
					}
				}
			}
			*s = res
		} else {
			var res []string
			res = make([]string, 0)
			for i := 0; i < len(s1); i++ {
				for k := 0; k < len(*s); k++ {
					//for j:=0; j<len(s[k]) ;j++ {
					res = append(res, strings.Join([]string{string(s1[i]), (*s)[k]}, ""))
					//}
				}
			}
			*s = res
		}
	}
}

//?
//func compute(s string , digit []int) {
//
//
//}
//
//Cannot use 'res' (type *[]int) as type []Type
//leetcode submit region end(Prohibit modification and deletion)
