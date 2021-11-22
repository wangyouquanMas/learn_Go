package main

import "fmt"

/*
   字典树创建
*/

var son = [10][5]int{}
var cnt = [10]int{}
var idx = 0

func main() {
	str := []byte{'a', 'b', 'c'}

	insert(str)

	fmt.Println("节点关系： ", son, "插入字符串个数：", cnt, "节点个数：", idx)
}

// 插入字符串形成字典树
func insert(str []byte) {
	p := 0
	for i := 0; i < len(str); i++ {
		u := str[i] - 'a'
		if son[p][u] == 0 {
			idx++
			son[p][u] = idx
		}
		p = son[p][u]
	}
	cnt[p]++
}

// 查询字符串出现的次数
func query(s []byte, son [][]int, cnt []int) int {
	next_pos := 0
	for i := 0; i < len(s); i++ {
		bias := s[i] - 'a'
		if son[next_pos][bias] == 0 {
			return 0
		}
		next_pos = son[next_pos][bias]
	}
	return cnt[next_pos]
}
