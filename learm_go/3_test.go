package learn_go

import (
	"fmt"
	"testing"
)

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
//
//
//
//示例 1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//提示：
//
//0 <= s.length <= 5 * 104
//s 由英文字母、数字、符号和空格组成

func Test3(T *testing.T) {
	fmt.Println(f3("abcabcbb"))
	fmt.Println(f3("pwwkew"))
	fmt.Println(f3("bbbbb"))
	fmt.Println(f3("a"))
	fmt.Println(f3("cdd"))
	fmt.Println(f3("acd"))
	fmt.Println(f3("aav"))

	fmt.Println(f3("abba"))
}

func f3(s string) int {

	cnt, r, l := 0, 0, 0
	flag := make(map[uint8]int)

	for {

		if l >= len(s) {
			break
		}

		if i, ok := flag[s[l]]; ok {
			if i >= r {
				r = i + 1
			}
		}
		flag[s[l]] = l
		cnt = f3Max(cnt, l-r+1)
		l++

	}

	if cnt == 0 {
		cnt = len(flag)
	}

	return cnt
}

func f3Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
