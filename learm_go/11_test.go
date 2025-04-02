package learn_go

import (
	"fmt"
	"testing"
)

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//返回容器可以储存的最大水量。
//
//说明：你不能倾斜容器。
//
//
//
//示例 1：
//
//
//
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7,6]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//示例 2：
//
//输入：height = [1,1]
//输出：1

func Test11(t *testing.T) {
	fmt.Println(f11([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(f11([]int{1, 1}))
}

// {

//   1:  [ 6,6 ]
//   8:  [ 1,6 ]

//}

func f11(height []int) int {

	l, r := 0, len(height)-1
	maxSum := 0
	for {
		if l == r {
			break
		}
		if height[l] < height[r] {
			maxSum = max(height[l]*(r-l), maxSum)
			l++
		} else {
			maxSum = max(height[r]*(r-l), maxSum)
			r--
		}
	}

	return maxSum
}

func f11_max(height []int) int {

	left, right := 0, 1

	sum := 0
	flag := 0
	for {
		if height[left] > height[right] {
			right++
		} else {
			sum += f11_2(height, left, right)
			left = right
			right++

		}

		if left == len(height)-1 || right == len(height) {
			sum += f11_2(height, left, right-1)
			break
		}
	}

	fmt.Println(flag)

	fmt.Println(sum)
	return sum
}

func f11_2(height []int, left, right int) int {
	sum := 0
	if height[left] >= height[right] {

		index := right - 1
		for {
			if left == index {
				break
			}

			if height[right] >= height[index] {
				sum += height[right] - height[index]
			}
			index--
		}

	} else {

		index := left + 1
		for {
			if index == right {
				break
			}
			if height[left] >= height[index] {
				sum += height[left] - height[index]

			}
			index++
		}

	}

	return sum
}
