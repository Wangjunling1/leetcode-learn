package a_learn

import (
	"testing"
)

//给你一个正整数 n。
//
//如果一个二进制字符串 x 的所有长度为 2 的
//子字符串
//中包含 至少 一个 "1"，则称 x 是一个 有效 字符串。
//
//返回所有长度为 n 的 有效 字符串，可以以任意顺序排列。
//示例 1：
//
//输入： n = 3
//
//输出： ["010","011","101","110","111"]
//
//解释：
//
//长度为 3 的有效字符串有："010"、"011"、"101"、"110" 和 "111"。
//
//示例 2：
//输入： n = 1
//
//输出： ["0","1"]
//
//解释：
//
//长度为 1 的有效字符串有："0" 和 "1"。
//提示：
//
//1 <= n <= 18

func Test3211(t *testing.T) {
	t.Log(f3211(3))
	t.Log(f3211(1))
}

func f3211(n int) []string {

	ans := make([]string, 0)
	strMap := make(map[string]struct{}, 0)
	dfs3211(n, "0", strMap)
	dfs3211(n, "1", strMap)
	for k, _ := range strMap {
		ans = append(ans, k)
	}
	return ans
}

func dfs3211(maxLen int, str string, strMap map[string]struct{}) {

	if len(str) >= maxLen {
		strMap[str] = struct{}{}
		return
	}

	if str[len(str)-1] != '0' {
		dfs3211(maxLen, str+"0", strMap)
		dfs3211(maxLen, str+"1", strMap)
	} else {
		dfs3211(maxLen, str+"1", strMap)
	}
}
