package learn_go

import (
	"fmt"
	"math"
	"testing"
)

//给你两个 正整数 l 和 r。对于任何数字 x，x 的所有正因数（除了 x 本身）被称为 x 的 真因数。
//
//如果一个数字恰好仅有两个 真因数，则称该数字为 特殊数字。例如：
//
//数字 4 是 特殊数字，因为它的真因数为 1 和 2。
//数字 6 不是 特殊数字，因为它的真因数为 1、2 和 3。
//返回区间 [l, r] 内 不是 特殊数字 的数字数量。
//
//
//
//示例 1：
//
//输入： l = 5, r = 7
//
//输出： 3
//
//解释：
//
//区间 [5, 7] 内不存在特殊数字。
//
//示例 2：
//
//输入： l = 4, r = 16
//
//输出： 11
//
//解释：
//
//区间 [4, 16] 内的特殊数字为 4 和 9。
//
//
//
//提示：
//
//1 <= l <= r <= 109

func Test3233(t *testing.T) {
	fmt.Println(f3233(4, 16))
	fmt.Println(f3233(5, 7))
	fmt.Println(findFactors(6))
}
func f3233(l, r int) int {

	n := int(math.Sqrt(float64(r)))
	v := make([]int, n+1)
	res := r - l + 1
	for i := 2; i <= n; i++ {
		if v[i] == 0 {
			if i*i >= l && i*i <= r {
				res--
			}
			for j := i * 2; j <= n; j += i {
				v[j] = 1
			}
		}
	}
	return res

}

func f3233BaoLi(l, r int) int {

	ans := 0

	for i := l; i <= r; i++ {
		if !findFactors(i) {
			ans++
		}
	}

	return ans
}

func findFactors(n int) bool {
	if n <= 0 {
		return false // 非正整数没有正因数
	}

	factors := 0
	for i := 1; i < n; i++ {
		if n%i == 0 { // 如果 i 能整除 n
			factors++
			if factors > 2 { // 如果真因数超过两个，则不是特殊数字
				return false
			}
		}
	}

	if factors == 2 { // 如果真因数恰好为两个，则是特殊数字
		return true
	}
	return false
}
