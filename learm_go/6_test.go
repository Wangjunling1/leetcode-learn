package learn_go

import (
	"fmt"
	"testing"
)

//将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
//比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//
//请你实现这个将字符串进行指定行数变换的函数：
//
//string convert(string s, int numRows);
//
//
//示例 1：
//
//输入：s = "PAYPALISHIRING", numRows = 3
//输出："PAHNAPLSIIGYIR"
//输出："PAHNAPLSIIGYIR"
//示例 2：
//输入：s = "PAYPALISHIRING", numRows = 4
//输出："PINALSIGYAHRPI"
//     "PINALSIGYAHRPI"
//解释：
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
//示例 3：
//
//输入：s = "A", numRows = 1
//输出："A"
//
//P
//A
//

//提示：
//
//1 <= s.length <= 1000
//s 由英文字母（小写和大写）、',' 和 '.' 组成
//1 <= numRows <= 1000

// PAYPALISHIRING

//  2
//P Y A I H I I G
//A P L S I R N
//P -> Y 2a

//  3
//P   A   H   N
//A P L S I I G
//Y   I   R
//P -> A 4

// 4
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
//P -> I 6

//   5
// P       H
// A     S I
// Y   I   R
// P L     I  G
// A       N

// P -> H 8
// A -> S 6 ,A  -> I 8
// PHASI YIR PLIG AN
// PHASI YIR PLIG AN
func Test6(T *testing.T) {
	fmt.Println(f6("PAYPALISHIRING", 4))
	fmt.Println(f6("AS", 1))
}

func f6(s string, numRows int) string {
	res := make([]byte, 0, len(s))

	if numRows == 1 {
		return s
	}
	skipRange := 2*numRows - 2
	skipCnt := 0
	for i := 0; i < numRows; {

		if skipCnt*skipRange+i > len(s)-1 {
			i++
			skipCnt = 0
			continue
		}

		res = append(res, s[skipCnt*skipRange+i])
		skipCnt++

		if skipCnt > 0 && skipCnt*skipRange-i <= len(s)-1 && i != 0 && i != numRows-1 {
			res = append(res, s[skipCnt*skipRange-i])
		}

	}
	return string(res)
}
