package learn_go

import (
	"fmt"
	"testing"
)

//给你一个 m x n 的二进制矩阵 grid 。
//
//如果矩阵中一行或者一列从前往后与从后往前读是一样的，那么我们称这一行或者这一列是 回文 的。
//
//你可以将 grid 中任意格子的值 翻转 ，也就是将格子里的值从 0 变成 1 ，或者从 1 变成 0 。
//
//请你返回 最少 翻转次数，使得矩阵 要么 所有行是 回文的 ，要么所有列是 回文的 。
//
//
//
//示例 1：
//
//输入：grid = [[1,0,0],[0,0,0],[0,0,1]]
//
//输出：2
//
//解释：
//
//
//
//将高亮的格子翻转，得到所有行都是回文的。
//
//示例 2：
//
//输入：grid = [[0,1],[0,1],[0,0]]
//
//输出：1
//
//解释：
//
//
//
//将高亮的格子翻转，得到所有列都是回文的。
//
//示例 3：
//
//输入：grid = [[1],[0]]
//
//输出：0
//
//解释：
//
//所有行已经是回文的。

func Test3239(t *testing.T) {
	fmt.Println(F3239([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}))
	fmt.Println(F3239([][]int{{0, 1}, {0, 1}, {0, 0}}))
	fmt.Println(F3239([][]int{{1}, {0}}))
}

func F3239(grid [][]int) int {

	cns := 0
	for i := range grid {
		jLen := len(grid[i]) - 1
		for j := 0; j <= jLen; {
			if grid[i][j] != grid[i][jLen] {
				cns++
			}
			j++
			jLen--
		}
	}
	cns1 := 0
	for i := range grid[0] {
		jLen := len(grid) - 1
		for j := 0; j <= jLen; {
			if grid[j][i] != grid[jLen][i] {
				cns1++
			}
			j++
			jLen--
		}
	}
	return mix3239(cns, cns1)
}

func mix3239(a, b int) int {
	if a > b {
		return b
	}
	return a
}
