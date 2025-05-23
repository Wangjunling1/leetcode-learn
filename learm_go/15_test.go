package learn_go

import (
	"fmt"
	"sort"
	"testing"
)

//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//
//
//
//
//示例 1：
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。
//示例 2：
//
//输入：nums = [0,1,1]
//输出：[]
//解释：唯一可能的三元组和不为 0 。
//示例 3：
//
//输入：nums = [0,0,0]
//输出：[[0,0,0]]
//解释：唯一可能的三元组和为 0 。
//
//
//提示：
//
//3 <= nums.length <= 3000
//-105 <= nums[i] <= 105

func Test15(t *testing.T) {
	//fmt.Println(f15([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(f15([]int{-2, 0, 1, 1, 2}))
}

func f15(nums []int) [][]int {
	sort.Ints(nums)

	list := make([][]int, 0)
	listMap := make(map[string][]int)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for {
			if left >= right {
				break
			}
			if nums[i]+nums[left]+nums[right] == 0 {
				listMap[fmt.Sprintf("%d,%d,%d", nums[i], nums[left], nums[right])] = []int{nums[i], nums[left], nums[right]}
				left++
			} else {
				if nums[i]+nums[left]+nums[right] < 0 {
					left++
				} else {
					right--
				}
			}

		}

	}
	for _, v := range listMap {
		list = append(list, v)
	}
	return list
}
