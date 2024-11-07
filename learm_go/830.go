package learn_go

import "fmt"

func main() {
	li := make([][]int, 0)
	s := "abcdddeeeeaabbbcd"
	start := 0
	end := 0
	var flag int32
	for i, v := range s {
		if flag != v {
			if i-start >= 3 {
				li = append(li, []int{start, i - 1})
			}
			flag = v
			start = i
		} else {
			end = i
		}
	}
	if end-start >= 2 {
		li = append(li, []int{start, end})
	}
	fmt.Print(li)
}
