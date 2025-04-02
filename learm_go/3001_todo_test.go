package learn_go

import "testing"

//现有一个下标从 1 开始的 8 x 8 棋盘，上面有 3 枚棋子。
//
//给你 6 个整数 a 、b 、c 、d 、e 和 f ，其中：
//
//(a, b) 表示白色车的位置。
//(c, d) 表示白色象的位置。
//(e, f) 表示黑皇后的位置。
//假定你只能移动白色棋子，返回捕获黑皇后所需的最少移动次数。
//
//请注意：
//
//车可以向垂直或水平方向移动任意数量的格子，但不能跳过其他棋子。
//象可以沿对角线方向移动任意数量的格子，但不能跳过其他棋子。
//如果车或象能移向皇后所在的格子，则认为它们可以捕获皇后。
//皇后不能移动。

func Test3001(t *testing.T) {

}

func F3001(a, b, c, d, e, f int) {
	// 移动 ab
	//
	// 边界 d,e
	// a <= e
	// b <= f
	//  if

	// 移动 cd
	// c <= e
	// d <= f
}

func MoveFunc(x, y, e, f, num int) {
	switch {
	case x < e:
		move()
	case x > e:

	case y < f:

	case y > f:

	}
}

func move(direction int, x, y int) (int, int) {
	switch direction {
	case 1: // 上
		return x, y + 1
	case 2: // 右上
		return x + 1, y + 1
	case 3: // 右
		return x + 1, y
	case 4: // 右下
		return x + 1, y - 1
	case 5: // 下
		return x, y - 1
	case 6: // 左下
		return x - 1, y - 1
	case 7: // 左
		return x - 1, y
	case 8: // 左上
		return x - 1, y + 1
	}
	return x, y
}
