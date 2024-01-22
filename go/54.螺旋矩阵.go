/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	row, column := len(matrix), len(matrix[0])
	if row == 0 || column == 0 {
		return []int{}
	}
	totalNum := row * column
	result := make([]int, totalNum)
	top := 0
	bottom := row - 1
	left := 0
	right := column - 1
	k := 0
	for k < totalNum {
		for i := left; i <= right; i, k = i+1, k+1 {
			result[k] = matrix[top][i]
		}
		if k == totalNum {
			break
		}
		top++
		for i := top; i <= bottom; i, k = i+1, k+1 {
			result[k] = matrix[i][right]
		}
		if k == totalNum {
			break
		}
		right--
		for i := right; i >= left; i, k = i-1, k+1 {
			result[k] = matrix[bottom][i]
		}
		if k == totalNum {
			break
		}
		bottom--
		for i := bottom; i >= top; i, k = i-1, k+1 {
			result[k] = matrix[i][left]
		}
		left++
	}
	return result
}

// @lc code=end
