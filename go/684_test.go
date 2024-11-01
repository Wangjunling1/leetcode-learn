package a_learn

import (
	"fmt"
	"testing"
)

//树可以看成是一个连通且 无环 的 无向 图。
//
//给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间，且这条附加的边不属于树中已存在的边。图的信息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。
//
//请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的那个。
//

//示例 1：
//输入: edges = [[1,2], [1,3], [2,3]]
//输出: [2,3]
//示例 2：
//输入: edges = [[1,2], [2,3], [3,4], [1,4], [1,5]]
//输出: [1,4]
//提示:
//
//n == edges.length
//3 <= n <= 1000
//edges[i].length == 2
//1 <= ai < bi <= edges.length
//ai != bi
//edges 中无重复元素
//给定的图是连通的

func Test684(t *testing.T) {
	edges := [][]int{{1, 2}, {2, 3}, {1, 4}, {3, 4}, {1, 5}}
	fmt.Println(f684(edges))
}

func f684(edges [][]int) []int {
	var ansInx []int
	uf := NewUnionFind(len(edges) + 1)
	for i := range edges {
		if uf.Connected(edges[i][0], edges[i][1]) {
			ansInx = edges[i]
		}
		uf.Union(edges[i][0], edges[i][1])

	}
	return ansInx
}

// UnionFind 结构体
type UnionFind struct {
	parent []int
	rank   []int
}

// NewUnionFind 创建并查集
func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank}
}

// Find 查找根节点
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

// Union 合并两个集合
func (uf *UnionFind) Union(x int, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		// 按秩合并
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

// Connected 判断两个元素是否在同一集合中
func (uf *UnionFind) Connected(x int, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
