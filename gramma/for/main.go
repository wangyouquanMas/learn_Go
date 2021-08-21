package main

import (
	"fmt"
	"sort"
)

func main() {
	test()
}

func test() {
	for i := 0; i < 100; i++ {
		if i > 0 {
			fmt.Println(i)
		}
	}
}

func threeSum(nums []int) [][]int {

	//
	var res [][]int
	res = make([][]int, 0)

	n := len(nums)
	//排序
	sort.Ints(nums)

	//var i int
	for i := 0; i < n; i++ {

		//if 不要括号， 需要中括号
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := n - 1
		for j := i + 1; j < k; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			//寻找>=0的最小的一个数 ，如果下一个数满足就用下一个数。 【避免重复】
			for j < k-1 && nums[i]+nums[j]+nums[k-1] >= 0 {
				k--
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[j], nums[i], nums[k]})
			}
		}
		11
	}
	return res
}
