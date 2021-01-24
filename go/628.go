package main

import (
	"fmt"
	"math"
)
func main() {
	numss := [] int{1,-1,-2,-3,2,4}
	max1,max2,max3 := math.MinInt64,math.MinInt64,math.MinInt64
	min1,min2 := math.MaxInt64,math.MaxInt64
	for _,i:=range numss{
		if i>max1{
			if i>max3 {
				max1, max2,max3 = max2, max3,i
			}else if i>max2{
				max1,max2=max2,i
			}else {
				max1=i
			}

		}
		if i<min2{
			if min1>i{
				min1,min2=i,min1
			}else {
				min2=i
			}
		}
	}
	if max1*max2*max3>min1*min2*max3{
		fmt.Println(max1*max2*max3)
	}else {
		fmt.Println(min1*min2*max3)
	}

}
