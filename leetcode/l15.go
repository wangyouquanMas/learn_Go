package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(input))
}

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {

	if nums == nil || len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)

	//var res [][]int
	//res = make([][]int,0)

	res := [][]int{}

	n := len(nums) - 1
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
				continue
			}

			if sum > 0 {
				k--
				continue
			}

			for k > j && nums[k] == nums[k-1] {
				k--
			}

			for k > j && nums[j] == nums[j+1] {
				j++
			}

			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				k-- //执行完之后，nums[j],nums[k]，nums[i] 在j<k情况下还有可能为0
				continue
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
