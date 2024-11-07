package learn_go

import (
	"fmt"
	"testing"
)

//14. 最长公共前缀
//已解答
//简单
//相关标签
//相关企业
//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//
//
//示例 1：
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//示例 2：
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//提示：
//
//1 <= strs.length <= 200
//0 <= strs[i].length <= 200
//strs[i] 仅由小写英文字母组成

func Test14(t *testing.T) {
	fmt.Println(f14([]string{"flower", "flow", "flight"}))
	fmt.Println(f14([]string{"dog", "racecar", "car"}))
	fmt.Println(f14([]string{"a"}))
}

func f14(nums []string) (ans string) {

	prMap := make(map[string]int)

	for _, v := range nums {
		for i := 0; i <= len(v); i++ {
			prMap[v[:i]]++
		}
	}
	fmt.Println(prMap)
	for k, v := range prMap {
		if v == len(nums) {
			if len(ans) < len(k) {
				ans = k
			}
		}
	}
	return
}
