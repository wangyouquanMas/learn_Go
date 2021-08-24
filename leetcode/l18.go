package main

import (
	"fmt"
	"sort"
)

// 1
2
func main() {

	input := []int{0, 0, 0, 0}

	fmt.Println(fourSum(input, 1))

}

func fourSum(nums []int, target int) [][]int {

	if nums == nil || len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)
	res := [][]int{}
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k := j + 1
			l := len(nums) - 1

			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum < target {
					k++
					continue
				}

				if sum > target {
					l--
					continue
				}

				for l > k && nums[l] == nums[l-1] {
					l--
				}

				for l > k && nums[k] == nums[k+1] {
					k++
				}

				res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
				l--
				k++
			}
		}

	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
