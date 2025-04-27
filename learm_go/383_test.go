package learn_go

import "testing"

//给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
//
//如果可以，返回 true ；否则返回 false 。
//
//magazine 中的每个字符只能在 ransomNote 中使用一次。
//
//
//
//示例 1：
//
//输入：ransomNote = "a", magazine = "b"
//输出：false
//示例 2：
//
//输入：ransomNote = "aa", magazine = "ab"
//输出：false
//示例 3：
//
//输入：ransomNote = "aa", magazine = "aab"
//输出：true
//
//
//提示：
//
//1 <= ransomNote.length, magazine.length <= 105
//ransomNote 和 magazine 由小写英文字母组成

func Test383(t *testing.T) {
	t.Log(f383("aa", "ab"))
	t.Log(f383("aa", "aab"))
	t.Log(f383("aa", "aba"))
}

func f383(ransomNote, magazine string) bool {

	map1 := [1000]int{}
	for _, v := range magazine {
		map1[v]++
	}

	for _, v := range ransomNote {
		map1[v]--
		if map1[v] < 0 {
			return false
		}
	}

	return true
}
