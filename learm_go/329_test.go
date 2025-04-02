package learn_go

import (
	"testing"
)

//给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
//字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
//
//进阶：
//
//如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
//

// 输入：s = "abc", t = "ahbgdc"
// 输出：true
// 示例 2：
//
// 输入：s = "axc", t = "ahbgdc"
// 输出：false
//
// 提示：
//
// 0 <= s.length <= 100
// 0 <= t.length <= 10^4
// 两个字符串都只由小写字符组成。
func Test329(t *testing.T) {

	t.Log(isSubStrDP("abc", "abcdeaac"))
	t.Log(isSubStrDP("acc", "ahbgdczzzzzzzzZ"))
}

// a b c d e f g h i j
//[0 1 2 3 4 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8]
//[5 1 2 3 4 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8]
//[5 8 2 3 4 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8]
//[5 8 7 3 4 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8]
//[5 8 7 8 4 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8]
//[5 8 7 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8]
//[6 8 7 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8]
//[8 8 7 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8]
//[8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8 8]

func isSubStr(s, t string) bool {
	sLen, tLen := len(s), len(t)
	sNum := 0
	for i := 0; i < tLen && sNum < sLen; i++ {
		if s[sNum] == t[i] {
			sNum++
		}
	}
	return sNum == len(s)
}

func isSubStrDP(s, t string) bool {
	n, m := len(s), len(t)
	f := make([][26]int, m+1)
	for i := 0; i < 26; i++ {
		f[m][i] = m
	}
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == byte(j+'a') {
				f[i][j] = i
			} else {
				f[i][j] = f[i+1][j]
			}
		}
	}
	add := 0
	for i := 0; i < n; i++ {
		if f[add][int(s[i]-'a')] == m {
			return false
		}
		add = f[add][int(s[i]-'a')] + 1
	}

	return true
}
