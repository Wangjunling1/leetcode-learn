package learn_go

//给你一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 k ，请你返回满足 0 <= i < j < n ，nums[i] == nums[j] 且 (i * j) 能被 k 整除的数对 (i, j) 的 数目 。
//
//
//
//示例 1：
//
//输入：nums = [3,1,2,2,2,1,3], k = 2
//输出：4
//解释：
//总共有 4 对数符合所有要求：
//- nums[0] == nums[6] 且 0 * 6 == 0 ，能被 2 整除。
//- nums[2] == nums[3] 且 2 * 3 == 6 ，能被 2 整除。
//- nums[2] == nums[4] 且 2 * 4 == 8 ，能被 2 整除。
//- nums[3] == nums[4] 且 3 * 4 == 12 ，能被 2 整除。
//示例 2：
//
//输入：nums = [1,2,3,4], k = 1
//输出：0
//解释：由于数组中没有重复数值，所以没有数对 (i,j) 符合所有要求。
//
//
//提示：
//
//1 <= nums.length <= 100
//1 <= nums[i], k <= 100
import "testing"

func Test2176(t *testing.T) {
	t.Log(f2176([]int{3, 1, 2, 2, 2, 1, 3}, 2))
}
func f2176(nums []int, k int) int64 {
	cnt := int64(0)
	numsLen := len(nums)
	for i := 0; i < numsLen-1; i++ {
		for j := i + 1; j < numsLen; j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				cnt++
			}
		}
	}
	return cnt
}
