package learn_go

import (
	"fmt"
	"testing"
)

// 给你两个整数 red 和 blue，分别表示红色球和蓝色球的数量。你需要使用这些球来组成一个三角形，满足第 1 行有 1 个球，第 2 行有 2 个球，第 3 行有 3 个球，依此类推。
//
// 每一行的球必须是 相同 颜色，且相邻行的颜色必须 不同。
//
// 返回可以实现的三角形的 最大 高度。
func Test3200(t *testing.T) {
	//fmt.Println("层：", func3200(2, 1))
	//fmt.Println("层：", func3200(10, 1))
	fmt.Println("层：", func3200(2, 4))
	fmt.Println("层：", func3200(4, 2))
	//fmt.Println("层：", func3200(2, 1))
	//fmt.Println("层：", func3200(10, 10))
}

func Max3200(a, b int) int {
	if a > b {
		return a
	}
	return b

}
func func3200(red int, blue int) int {

	level, co1, co2 := 0, 0, 0
	for {
		level++
		if level%2 == 0 {
			co2 += level
		} else {
			co1 += level
		}
		if co1 > red || co2 > blue {
			return level - 1
		}
	}
}
