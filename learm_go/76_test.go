package learn_go

import (
	"fmt"
	"math"
	"testing"
)

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//
//
//注意：
//
//对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
//如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//示例 1：
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//示例 2：
//
//输入：s = "a", t = "a"
//输出："a"
//解释：整个字符串 s 是最小覆盖子串。
//示例 3:
//
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。
//
//
//提示：
//
//m == s.length
//n == t.length
//1 <= m, n <= 105
//s 和 t 由英文字母组成

func Test76(T *testing.T) {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("a", "b"))
	fmt.Println(minWindow("abc", "ac"))
}

func f76(s, t string) string {
	if len(s) < len(t) {
		return ""
	}

	atMap := map[byte]int{}
	for i := range t {
		atMap[t[i]]++
	}

	ans := s + t
	for i := 0; i < len(s); i++ {

		if _, ok := atMap[s[i]]; !ok {
			continue
		}

		tMap := map[byte]int{}
		for k, v := range atMap {
			tMap[k] = v
		}

		for j := i; j < len(s); j++ {

			if tMap[s[j]] > 0 {
				tMap[s[j]]--
				if tMap[s[j]] == 0 {
					delete(tMap, s[j])
				}
			}

			if len(tMap) == 0 {
				ans = min76(ans, s[i:j+1])
				break
			}

		}
	}

	if ans == s+t {
		return ""
	}

	return ans
}

func min76(a, b string) string {
	if len(a) > len(b) {
		return b
	}
	return a
}

func minWindow(s string, t string) string {
	cnt := [128]int{}
	diff := 0
	for _, x := range t {
		cnt[x]++
		if cnt[x] == 1 {
			diff++
		}
	}

	left, l := 0, math.MaxInt
	resL, resR := -1, -1

	for right, x := range s {
		cnt[x]--
		if cnt[x] == 0 {
			diff--
		}

		for diff == 0 {
			if l > right-left+1 {
				l = right - left + 1
				resL, resR = left, right
			}

			tmp := s[left]
			left++
			cnt[tmp]++
			if cnt[tmp] == 1 {
				diff++
			}
		}
	}
	if resL == -1 {
		return ""
	} else {
		return s[resL : resR+1]
	}
}
