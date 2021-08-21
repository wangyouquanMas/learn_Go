//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和
//。假定每组输入只存在唯一答案。
//
//
//
// 示例：
//
// 输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 10^3
// -10^3 <= nums[i] <= 10^3
// -10^4 <= target <= 10^4
//
// Related Topics 数组 双指针 排序 👍 856 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"fmt"
	"sort"
)

func main() {

	var nums = []int{1, 1, 1, 1}
	target := 1
	res := threeSumClosest(nums, target)
	fmt.Println(res)
}

func threeSumClosest(nums []int, target int) int {

	if nums == nil || len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)
	k := n - 1
	var min int = 1e8
	var sum2 int
	for i := 0; i < n; i++ {
		for j := i + 1; j < k; j++ {

			for j < k-1 && nums[i]+nums[j]+nums[k-1] >= target {
				k--
			}

			sum1 := nums[i] + nums[j] + nums[k]
			if j < k-1 {
				sum2 = nums[i] + nums[j] + nums[k-1]
			}
			res := compare(sum1, sum2, target)
			min = compare(min, res, target)
		}
	}
	return min
}
func compare(a, b, c int) int {
	if abs(a-c) > abs(c-b) {
		return b
	} else {
		return a
	}
	return 0
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
