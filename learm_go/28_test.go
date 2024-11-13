package learn_go

import (
	"testing"
)

//给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
//
//
//
//示例 1：
//
//输入：haystack = "sadbutsad", needle = "sad"
//输出：0
//解释："sad" 在下标 0 和 6 处匹配。
//第一个匹配项的下标是 0 ，所以返回 0 。
//示例 2：
//
//输入：haystack = "leetcode", needle = "leeto"
//输出：-1
//解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
//
//
//提示：
//
//1 <= haystack.length, needle.length <= 104
//haystack 和 needle 仅由小写英文字符组成

// 如何实现kmp TODO

func Test28(t *testing.T) {
	//t.Log(F28("sadbutsad", "sad"))
	//t.Log(F28("leetcode", "leeto"))
	////t.Log(F28("gello", "ll"))
	//t.Log(F28("ba", "a"))
	//t.Log(F28("a", "a"))
	t.Log(F28("mississippi", "issip"))
}

func F28(haystack, needle string) int {
	st, j, n, m := 0, 0, len(haystack), len(needle)
	for st+j <= n {

		if j == m {
			return st
		}

		if st+j < n && haystack[st+j] == needle[j] {
			j++
		} else {
			st += 1
			j = 0
		}

	}
	return -1
}
