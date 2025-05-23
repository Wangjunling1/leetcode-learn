package learn_go

import (
	"fmt"
	"strings"
	"testing"
)

//给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
//
//你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
//
//要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
//
//文本的最后一行应为左对齐，且单词之间不插入额外的空格。
//
//注意:
//
//单词是指由非空格字符组成的字符序列。
//每个单词的长度大于 0，小于等于 maxWidth。
//输入单词数组 words 至少包含一个单词。
//
//
//示例 1:
//
//输入: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
//输出:
//[
//   "This    is    an",
//   "example  of text",
//   "justification.  "
//]
//示例 2:
//
//输入:words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
//输出:
//[
//  "What   must   be",
//  "acknowledgment  ",
//  "shall be        "
//]
//解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
//     因为最后一行应为左对齐，而不是左右两端对齐。
//     第二行同样为左对齐，这是因为这行只包含一个单词。
//示例 3:
//
//输入:words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]，maxWidth = 20
//输出:
//[
//  "Science  is  what we",
//   Science  is  what we
//  "understand      well",
//  "enough to explain to",
//  "a  computer.  Art is",
//  "everything  else  we",
//  "do                  "
//]
//
//
//提示:
//
//1 <= words.length <= 300
//1 <= words[i].length <= 20
//words[i] 由小写英文字母和符号组成
//1 <= maxWidth <= 100
//words[i].length <= maxWidth

func Test68(t *testing.T) {

	//fmt.Println(9 / 3)

	//fmt.Println(Aligning68("Science is what", 20, 4))
	//fmt.Println(len(Aligning68("Science is what", 20, 4)))
	//fmt.Println(Aligning68("a computer. Art is  ", 20, 4))
	//fmt.Println(Aligning68("understand well", 20, 4))
	//fmt.Println(Aligning68("     understand well", 20, 3))

	fmt.Println(F68([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20))
	fmt.Println(F68([]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16))
	fmt.Println(F68([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))

}
func F68(words []string, maxWidth int) []string {

	cnt := ""
	res := make([]string, 0)
	for i := range words {
		if len(cnt+words[i]) > maxWidth {
			res = append(res, Aligning68(cnt, maxWidth, 4))
			cnt = ""
		}
		cnt += words[i] + " "
	}
	if cnt != "" {
		res = append(res, Aligning68(cnt, maxWidth, 3))
	}
	return res
}

// Aligning68 对齐 direction: 1 右对齐 2 居中对齐 3 左对齐 4: 两端对齐
func Aligning68(s string, maxWidth, direction int) string {
	sLen := make([]string, 0)
	for _, v := range strings.Split(s, " ") {
		if v == "" {
			continue
		}
		sLen = append(sLen, v)
	}
	sum := maxWidth - len(strings.Join(sLen, ""))
	wordLen := len(sLen) - 1

	if wordLen == 0 {
		direction = 3
	}

	ans := ""
	switch direction {
	case 1: // 右对齐
	case 2: // 居中对齐
	case 3: // 左对齐

		for i := range sLen {
			ans += sLen[i]
			if len(ans) == maxWidth {
				continue
			}
			ans += " "
		}

		ans += strings.Repeat(" ", maxWidth-len(ans))
		return ans

	case 4: // 两端对齐
		for i := range sLen {
			ans += sLen[i]
			if len(ans) == maxWidth {
				continue
			}
			ans += strings.Repeat(" ", sum/wordLen)
			if sum%wordLen > 0 {
				ans += " "
				sum--
			}
		}
		return ans

	}
	return s
}
