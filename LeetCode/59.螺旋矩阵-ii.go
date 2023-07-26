/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start

// Official方法一：模拟
// type pair struct {
// 	x int
// 	y int
// }
// // 枚举方向对应的xy增量
// var dirs = []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} //右下左上
// func generateMatrix(n int) [][]int {
// 	// n*n的切片
// 	matrix := make([][]int, n)
// 	for i := range matrix {
// 		matrix[i] = make([]int, n)
// 	}

// 	row, col, dirIdx := 0, 0, 0
// 	for i := 1; i <= n*n; i++ {
// 		matrix[row][col] = i
// 		dir := dirs[dirIdx]
// 		// 判断是否超出矩阵边界
// 		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
// 			dirIdx = (dirIdx + 1) % 4 //顺时针旋转至下一方向
// 			dir = dirs[dirIdx]
// 		}
// 		row += dir.x
// 		col += dir.y
// 	}
// 	return matrix
// }

// Official方法二：按层模拟
// func generateMatrix(n int) [][]int {
// 	matrix := make([][]int, n)
// 	for i := range matrix {
// 		matrix[i] = make([]int, n)
// 	}

// 	num := 1
// 	left, right, top, bottom := 0, n-1, 0, n-1
// 	for left <= right && top <= bottom {
// 		for column := left; column <= right; column++ {
// 			matrix[top][column] = num
// 			num++
// 		}
// 		for row := top + 1; row <= bottom; row++ {
// 			matrix[row][right] = num
// 			num++
// 		}
// 		if left < right && top < bottom {
// 			for column := right - 1; column > left; column-- {
// 				matrix[bottom][column] = num
// 				num++
// 			}
// 			for row := bottom; row > top; row-- {
// 				matrix[row][left] = num
// 				num++
// 			}
// 		}
// 		left++
// 		right--
// 		top++
// 		bottom--
// 	}
// 	return matrix
// }

// 模拟(逐层)
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	top := 0
	bottom := n - 1
	left := 0
	right := n - 1
	k := 1
	lens = n * n
	for k <= lens {
		for i := left; i <= right; i, k = i+1, k+1 {
			matrix[top][i] = k
		}
		top++
		for i := top; i <= bottom; i, k = i+1, k+1 {
			matrix[i][right] = k
		}
		right--
		for i := right; i >= left; i, k = i-1, k+1 {
			matrix[bottom][i] = k
		}
		bottom--
		for i := bottom; i >= top; i, k = i-1, k+1 {
			matrix[i][left] = k
		}
		left++
	}
	return matrix
}

// @lc code=end

