package learn_go

// 给你两个下标从 0 开始且长度为 n 的整数数组 nums1 和 nums2 ，两者都是 [0, 1, ..., n - 1] 的 排列 。
//
//好三元组 指的是 3 个 互不相同 的值，且它们在数组 nums1 和 nums2 中出现顺序保持一致。换句话说，如果我们将 pos1v 记为值 v 在 nums1 中出现的位置，pos2v 为值 v 在 nums2 中的位置，那么一个好三元组定义为 0 <= x, y, z <= n - 1 ，且 pos1x < pos1y < pos1z 和 pos2x < pos2y < pos2z 都成立的 (x, y, z) 。
//
//请你返回好三元组的 总数目 。
//
//
//
//示例 1：
//
//输入：nums1 = [2,0,1,3], nums2 = [0,1,2,3]
//输出：1
//解释：
//总共有 4 个三元组 (x,y,z) 满足 pos1x < pos1y < pos1z ，分别是 (2,0,1) ，(2,0,3) ，(2,1,3) 和 (0,1,3) 。
//这些三元组中，只有 (0,1,3) 满足 pos2x < pos2y < pos2z 。所以只有 1 个好三元组。
//示例 2：
//
//输入：nums1 = [4,0,1,3,2], nums2 = [4,1,0,2,3]
//输出：4
//解释：总共有 4 个好三元组 (4,0,3) ，(4,0,2) ，(4,1,3) 和 (4,1,2) 。
//
//
//提示：
//
//n == nums1.length == nums2.length
//3 <= n <= 105
//0 <= nums1[i], nums2[i] <= n - 1
//nums1 和 nums2 是 [0, 1, ..., n - 1] 的排列。

import (
	"fmt"
	"testing"
)

func Test2179(t *testing.T) {
	fmt.Println(f2179([]int{12, 2, 1, 10, 7, 6, 13, 9, 11, 4, 5, 3, 14, 15, 8, 0}, []int{9, 11, 10, 4, 5, 2, 1, 8, 7, 3, 12, 13, 6, 15, 14, 0}))
	t.Log(f2179([]int{4, 0, 1, 3, 2}, []int{4, 1, 0, 2, 3}))
	t.Log(f2179([]int{2, 0, 1, 3}, []int{0, 1, 2, 3}))
}

type FenwickTree struct {
	tree []int
}

func fenwickTree(size int) *FenwickTree {
	return &FenwickTree{tree: make([]int, size+1)}
}

func (ft *FenwickTree) update(index, delta int) {
	index++
	for index < len(ft.tree) {
		ft.tree[index] += delta
		index += index & -index
	}
}

func (ft *FenwickTree) query(index int) int {
	index++
	res := 0
	for index > 0 {
		res += ft.tree[index]
		index -= index & -index
	}
	return res
}

func f2179(nums1, nums2 []int) int64 {
	n := len(nums1)
	list1 := make([]int, n)
	reversedIndexMapping := make([]int, n)
	for i, v := range nums1 {
		list1[v] = i
	}
	for i, v := range nums2 {
		reversedIndexMapping[list1[v]] = i
	}

	fmt.Println(list1, reversedIndexMapping)

	tree := fenwickTree(n)
	var res int64
	for value := 0; value < n; value++ {
		pos := reversedIndexMapping[value]
		left := tree.query(pos)
		tree.update(pos, 1)
		right := (n - 1 - pos) - (value - left)
		res += int64(left * right)
	}

	return res
}
func f2179_N3(nums1, nums2 []int) int64 {

	list := [105]int{}
	for i, v := range nums2 {
		list[v] = i
	}

	numsLen := len(nums1)
	ans := int64(0)
	for i := 0; i < numsLen-2; i++ {
		for j := i + 1; j < numsLen-1; j++ {
			for k := j + 1; k < numsLen; k++ {
				if nums1[i] != nums1[j] && nums1[j] != nums1[k] && nums1[i] != nums1[k] {
					if list[nums1[i]] < list[nums1[j]] && list[nums1[j]] < list[nums1[k]] {
						ans++
					}
				}
			}
		}
	}
	return ans
}

func subF2179(nums []int, hao map[string]int, op string) {
	numsLen := len(nums)
	for i := 0; i < numsLen-2; i++ {
		for j := i + 1; j < numsLen-1; j++ {
			for k := j + 1; k < numsLen; k++ {
				if nums[i] != nums[j] && nums[j] != nums[k] && nums[i] != nums[k] {

					hao[fmt.Sprintf("%v:%v:%v", nums[i], nums[j], nums[k])]++

				}
			}
		}
	}
	return
}
