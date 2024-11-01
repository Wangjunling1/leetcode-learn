package a_learn

import (
	"fmt"
	"sort"
	"testing"
)

//给你一个仅由小写英文字母组成的字符串 s 。
//
//如果一个字符串仅由单一字符组成，那么它被称为 特殊 字符串。例如，字符串 "abc" 不是特殊字符串，而字符串 "ddd"、"zz" 和 "f" 是特殊字符串。
//
//返回在 s 中出现 至少三次 的 最长特殊子字符串 的长度，如果不存在出现至少三次的特殊子字符串，则返回 -1 。
//
//子字符串 是字符串中的一个连续 非空 字符序列。
//
//
//
//示例 1：
//
//输入：s = "aaaa"
//输出：2
//解释：出现三次的最长特殊子字符串是 "aa" ：子字符串 "aaaa"、"aaaa" 和 "aaaa"。
//可以证明最大长度是 2 。
//示例 2：
//
//输入：s = "abcdef"
//输出：-1
//解释：不存在出现至少三次的特殊子字符串。因此返回 -1 。
//示例 3：
//
//输入：s = "abcaba"
//输出：1
//解释：出现三次的最长特殊子字符串是 "a" ：子字符串 "abcaba"、"abcaba" 和 "abcaba"。
//可以证明最大长度是 1 。
//
//
//提示：
//
//3 <= s.length <= 5 * 105
//s 仅由小写英文字母组成。

// 1. l1 - 2
// 2. l1=l2 l1>l2  min(l1-1,l2)
// 3. l1
func Test2982(t *testing.T) {
	s := "abcdef"
	cntMap := make(map[string][]int, 0)
	cnt := 0
	ans := -1
	for i := 0; i < len(s); i++ {
		cnt += 1
		if i == len(s)-1 || s[i] != s[i+1] {
			cntMap[string(s[i])] = append(cntMap[string(s[i])], cnt)
			cnt = 0
		}
	}
	for _, v := range cntMap {
		v = append(v, 0, 0)
		sort.Ints(v)
		Reverse(v)

		l1, l2, l3 := v[0], v[1], v[2]
		ans = Max(ans, Max(l1-2, Min(l1-1, l2), l3))
		if ans == 0 {
			ans = -1
		}
	}
	//fmt.Println(ans, max(max(l1-2, min(l1-1, l2)), l3))
	fmt.Println(ans)
}

func Min[T int | int8 | int16 | int32 | int64 | float32 | float64 | uint | string](args ...T) (ans T) {
	ans = args[0]
	for _, a := range args {
		if a < ans {
			ans = a
		}
	}
	return
}

func Max[T int | int8 | int16 | int32 | int64 | float32 | float64 | uint | string](args ...T) (ans T) {
	ans = args[0]
	for _, a := range args {
		if a > ans {
			ans = a
		}
	}
	return
}

func Reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
