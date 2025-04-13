package learn_go

import (
	"fmt"
	"testing"
)

//`给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。
//
// s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。
//
//例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"， "cdefab"，"efabcd"， 和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。
//返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。
//
//
//
//示例 1：
//
//输入：s = "barfoothefoobarman", words = ["foo","bar"]
//输出：[0,9]
//解释：因为 words.length == 2 同时 words[i].length == 3，连接的子字符串的长度必须为 6。
//子串 "barfoo" 开始位置是 0。它是 words 中以 ["bar","foo"] 顺序排列的连接。
//子串 "foobar" 开始位置是 9。它是 words 中以 ["foo","bar"] 顺序排列的连接。
//输出顺序无关紧要。返回 [9,0] 也是可以的。
//示例 2：
//
//输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
//输出：[]
//解释：因为 words.length == 4 并且 words[i].length == 4，所以串联子串的长度必须为 16。
//s 中没有子串长度为 16 并且等于 words 的任何顺序排列的连接。
//所以我们返回一个空数组。
//示例 3：
//
//输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
//输出：[6,9,12]
//解释：因为 words.length == 3 并且 words[i].length == 3，所以串联子串的长度必须为 9。
//子串 "foobarthe" 开始位置是 6。它是 words 中以 ["foo","bar","the"] 顺序排列的连接。
//子串 "barthefoo" 开始位置是 9。它是 words 中以 ["bar","the","foo"] 顺序排列的连接。
//子串 "thefoobar" 开始位置是 12。它是 words 中以 ["the","foo","bar"] 顺序排列的连接。
//
//
//提示：
//
//1 <= s.length <= 104
//1 <= words.length <= 5000
//1 <= words[i].length <= 30
//words[i] 和 s 由小写英文字母组成

func Test30(T *testing.T) {
	fmt.Println(f30("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Println(f30("ddddddd", []string{"dd", "d", "dd"}))
	fmt.Println(f30("wordgoodgoodwordbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(f30("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(f30("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}

func f30(s string, words []string) (res []int) {
	//var res []int
	//for i := 0; i < len(s); i++ {
	//
	//	if dg30(s, i, words) {
	//		res = append(res, i)
	//	}
	//
	//}
	//fmt.Println(res)
	res = f30V2HuaKuai(s, words)

	return res

}

func f30V2HuaKuai(s string, words []string) []int {
	wordMap := make(map[string]int)
	res := make([]int, 0)
	wordLen := len(words[0])
	sumLen := 0
	for _, v := range words {
		wordMap[v]++
		sumLen += len(v)
	}
	for i := 0; i < len(s)-sumLen+1; i++ {
		tmpMap := make(map[string]int)
		for start := i; start < len(s)-wordLen+1; start += wordLen {

			word := s[start : start+wordLen]
			if _, ok := wordMap[word]; !ok {
				break
			}
			tmpMap[word]++
			if tmpMap[word] > wordMap[word] {
				break
			}

			if start-i+wordLen == sumLen {
				res = append(res, i)
				break
			}

		}

	}
	return res
}

func dg30(s string, start int, words []string) bool {

	for i, v := range words {

		if start+len(v) > len(s) {
			return false
		}
		if s[start:start+len(v)] == v {
			start += len(v)
			if 1 < len(words) {
				newWords := make([]string, 0)
				newWords = append(newWords, words[:i]...)
				newWords = append(newWords, words[i+1:]...)
				return dg30(s, start, newWords)
			} else {
				return true
			}

		}
	}
	return false
}
