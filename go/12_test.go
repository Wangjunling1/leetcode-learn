package a_learn

import (
	"fmt"
	"testing"
)

func Test12(t *testing.T) {
	//七个不同的符号代表罗马数字，其值如下：
	//
	//符号	值
	//I	1
	//V	5
	//X	10
	//L	50
	//C	100
	//D	500
	//M	1000
	//罗马数字是通过添加从最高到最低的小数位值的转换而形成的。将小数位值转换为罗马数字有以下规则：
	//
	//如果该值不是以 4 或 9 开头，请选择可以从输入中减去的最大值的符号，将该符号附加到结果，减去其值，然后将其余部分转换为罗马数字。
	//如果该值以 4 或 9 开头，使用 减法形式，表示从以下符号中减去一个符号，例如 4 是 5 (V) 减 1 (I): IV ，9 是 10 (X) 减 1 (I)：IX。仅使用以下减法形式：4 (IV)，9 (IX)，40 (XL)，90 (XC)，400 (CD) 和 900 (CM)。
	//只有 10 的次方（I, X, C, M）最多可以连续附加 3 次以代表 10 的倍数。你不能多次附加 5 (V)，50 (L) 或 500 (D)。如果需要将符号附加4次，请使用 减法形式。
	//给定一个整数，将其转换为罗马数字。
	//示例 1：
	//输入：num = 3749
	//输出： "MMMDCCXLIX"
	//解释：
	//
	//3000 = MMM 由于 1000 (M) + 1000 (M) + 1000 (M)
	// 700 = DCC 由于 500 (D) + 100 (C) + 100 (C)
	//  40 = XL 由于 50 (L) 减 10 (X)
	//   9 = IX 由于 10 (X) 减 1 (I)
	//注意：49 不是 50 (L) 减 1 (I) 因为转换是基于小数位
	//示例 2：
	//
	//输入：num = 58
	//
	//输出："LVIII"
	//
	//解释：
	//
	//50 = L
	// 8 = VIII
	//示例 3：
	//
	//输入：num = 1994
	//
	//输出："MCMXCIV"
	//
	//解释：
	//
	//1000 = M
	// 900 = CM
	//  90 = XC
	//   4 = IV
	//
	//
	//提示：
	//
	//1 <= num <= 3999

	fmt.Println(f12(1994))
}

var RomanMap = map[int]string{
	1:    "I",
	2:    "II",
	3:    "III",
	4:    "IV",
	5:    "V",
	6:    "VI",
	7:    "VII",
	8:    "VIII",
	9:    "IX",
	10:   "X",
	20:   "XX",
	30:   "XXX",
	40:   "XL",
	50:   "L",
	60:   "LX",
	70:   "LXX",
	80:   "LXXX",
	90:   "XC",
	100:  "C",
	200:  "CC",
	300:  "CCC",
	400:  "CD",
	500:  "D",
	600:  "DC",
	700:  "DCC",
	800:  "DCCC",
	900:  "CM",
	1000: "M",
	2000: "MM",
	3000: "MMM",
}

func f12(num int) (ans string) {
	//numStr := fmt.Sprintf("%d", num)
	//for i := len(numStr); i > 0; i-- {
	//	jl := num / int(math.Pow(10, float64(i-1)))
	//	c := jl * int(math.Pow(10, float64(i-1)))
	//	num -= c
	//	ans += RomanMap[c]
	//}
	n1 := (num / 1000) * 1000
	num -= n1
	n2 := (num / 100) * 100
	num -= n2
	n3 := (num / 10) * 10
	num -= n3
	n4 := num

	return RomanMap[n1] + RomanMap[n2] + RomanMap[n3] + RomanMap[n4]
}
