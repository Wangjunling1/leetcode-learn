package a_learn

import (
	"fmt"
	"testing"
)

/*
在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。

示例 1:

输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
输出: 3
解释:
从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
因此，3 可为起始索引。
示例 2:

sum = gas[i]
if sum  - cost[i] > 0 {
	sum = - cost[i] + gas[i+1]
}



*/

func Test134(t *testing.T) {
	gas := []int{5, 8, 2, 8}
	cost := []int{6, 5, 6, 6}
	//gas = []int{1, 2, 3, 4, 5}
	//cost = []int{3, 4, 5, 1, 2}
	//gas = []int{3, 1, 1}
	//cost = []int{1, 2, 2}
	gas = []int{4, 5, 2, 6, 5, 3}
	cost = []int{3, 2, 7, 3, 2, 9}

	sum := 0
	i, j, gasLen := 0, 0, len(gas)

	for i = 0; i < len(gas); i++ {
		sum = gas[i]
		j = i
		sumG := 0
		sumC := 0
		for {

			sumG += gas[j]
			sumC += cost[j]

			if sumG < sumC {
				i = j
				break
			}

			sum -= cost[j]
			if sum < 0 {
				fmt.Println("false", i)
				break
			}

			j = (j + 1) % gasLen
			fmt.Println(j)
			sum += gas[j]

			if i == j {
				fmt.Println("true", i)
				break
			}
		}

	}

	//for i, n := 0, len(gas); i < n; {
	//	sumOfGas, sumOfCost, cnt := 0, 0, 0
	//	for cnt < n {
	//		j := (i + cnt) % n
	//		sumOfGas += gas[j]
	//		sumOfCost += cost[j]
	//		if sumOfCost > sumOfGas {
	//			break
	//		}
	//		cnt++
	//	}
	//	if cnt == n {
	//		fmt.Println("牛逼的：", i)
	//		break
	//	} else {
	//		i += cnt + 1
	//	}
	//}

	fmt.Println(1 % 4)
	fmt.Println(122 % 4)
}
