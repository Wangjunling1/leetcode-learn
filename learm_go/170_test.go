package learn_go

// 在股票交易中，如果前一天的股价高于后一天的股价，则可以认为存在一个「交易逆序对」。请设计一个程序，输入一段时间内的股票交易记录 record，返回其中存在的「交易逆序对」总数。
//
//
//
//示例 1：
//
//输入：record = [9, 7, 5, 4, 6]
//输出：8
//解释：交易中的逆序对为 (9, 7), (9, 5), (9, 4), (9, 6), (7, 5), (7, 4), (7, 6), (5, 4)。
import "testing"

func Test170(t *testing.T) {
	t.Log(f170([]int{9, 7, 5, 4, 6}))
}

func f170(record []int) int {
	res := 0
	for i := 0; i < len(record); i++ {
		for j := i + 1; j < len(record); j++ {
			if record[i] > record[j] {
				res++
			}
		}
	}
	return res
}
