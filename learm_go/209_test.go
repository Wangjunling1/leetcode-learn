package learn_go

import (
	"fmt"
	"testing"
)

// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//
// 示例 1：
//
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
// 示例 2：
//
// 输入：target = 4, nums = [1,4,4]
// 输出：1
// 示例 3：
//
// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0
//
// 提示：
//
// 1 <= target <= 109
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 104
//
// 进阶：
//
// 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
func Test209(T *testing.T) {
	//fmt.Println(f209v2(7, []int{2, 3, 1, 2, 4, 3}))
	//fmt.Println(f209(4, []int{1, 4, 4}))
	//fmt.Println(f209v2(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(f209v2(213, []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}))
	fmt.Println(f209v2(15, []int{1, 2, 3, 4, 5}))
	fmt.Println(f209v2(6, []int{10, 2, 3}))
}
func f209(target int, nums []int) int {
	cnt := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				cnt = mix209(cnt, j-i+1)
				break
			}
		}

	}
	if cnt == len(nums)+1 {
		return 0
	}
	return cnt
}

func f209v2(target int, nums []int) int {
	cnt := len(nums) + 1

	i, j, sum := 0, 0, nums[0]
	for {
		if i >= len(nums) {
			break
		}

		if sum >= target {
			cnt = mix209(cnt, j-i+1)
		}

		// 指针移动
		if sum < target && j < len(nums)-1 {
			j++
			sum += nums[j]
		} else {
			sum -= nums[i]
			i++
		}

	}

	if cnt == len(nums)+1 {
		return 0
	}

	return cnt
}
func mix209(a, b int) int {
	if a > b {
		return b
	}
	return a
}
