package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(input))
}

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {

	if strs == nil || len(strs) == 0 {
		return nil
	}

	res := [][]string{}

	if len(strs) == 1 {
		res = append(res, strs)
		return res
	}

	var isShowMap = make(map[string][]string)

	for _, v := range strs {
		s := SortString(v)
		isShowMap[s] = append(isShowMap[s], v)
	}
	.
	for _, v := range isShowMap {
		res = append(res, v)
	}

	return res
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

//leetcode submit region end(Prohibit modification and deletion)
