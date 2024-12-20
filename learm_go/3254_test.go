package learn_go

import (
	"testing"
)

// 给你一个长度为 n 的整数数组 nums 和一个正整数 k 。
//
//一个数组的 能量值 定义为：
//
//如果 所有 元素都是依次 连续 且 上升 的，那么能量值为 最大 的元素。
//否则为 -1 。
//你需要求出 nums 中所有长度为 k 的子数组 的能量值。
//
//请你返回一个长度为 n - k + 1 的整数数组 results ，其中 results[i] 是子数组 nums[i..(i + k - 1)] 的能量值。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3,4,3,2,5], k = 3
//
//输出：[3,4,-1,-1,-1]
//
//解释：
//
//nums 中总共有 5 个长度为 3 的子数组：
//
//[1, 2, 3] 中最大元素为 3 。
//[2, 3, 4] 中最大元素为 4 。
//[3, 4, 3] 中元素 不是 连续的。
//[4, 3, 2] 中元素 不是 上升的。
//[3, 2, 5] 中元素 不是 连续的。
//示例 2：
//
//输入：nums = [2,2,2,2,2], k = 4
//
//输出：[-1,-1]
//
//示例 3：
//
//输入：nums = [3,2,3,2,3,2], k = 2
//
//输出：[-1,3,-1,3,-1]
//
//
//
//提示：
//
//1 <= n == nums.length <= 500
//1 <= nums[i] <= 105
//1 <= k <= n

func Test3254(t *testing.T) {
	//t.Log(F3254(3, []int{1, 2, 3, 4, 3, 2, 5}))
	//t.Log(F3254(4, []int{2, 2, 2, 2, 2}))
	//t.Log(F3254(2, []int{3, 2, 3, 2, 3, 2}))
	//t.Log(F3254(1, []int{1, 2, 3, 4, 5, 6, 7}))
	//t.Log(F3254(1, []int{1}))
	//t.Log(F3254(2, []int{2, 3}))

	t.Log(F3254V2(3, []int{1, 2, 3, 4, 3, 2, 5}))
	t.Log(F3254V2(4, []int{2, 2, 2, 2, 2}))
	t.Log(F3254V2(2, []int{3, 2, 3, 2, 3, 2}))
	t.Log(F3254V2(1, []int{1, 2, 3, 4, 5, 6, 7}))
	t.Log(F3254V2(1, []int{1, 4}))

}

func F3254(k int, nums []int) []int {
	i, j, n := 0, 0, len(nums)
	for i < n {

		if j-i+1 == k {
			nums[i] = nums[j]
			i++
			j = i
			continue
		}

		if j >= n-1 {
			nums[i] = -1
			i++
			j = i
			continue
		}

		if nums[j]+1 == nums[j+1] {
			j++
		} else {
			nums[i] = -1
			i++
			j = i
		}
	}

	return nums[:n-k+1]
}

// TODO 优化 算法有问题 没有时间
func F3254V2(k int, nums []int) []int {
	cnt := 0
	idx := 0
	n := len(nums)
	for i, v := range nums {
		if i == 0 || v != nums[i-1]+1 {
			cnt = 1
		} else {
			cnt++
		}

		if cnt >= k {
			nums[idx] = v
			idx++
		}
	}
	return nums
}
