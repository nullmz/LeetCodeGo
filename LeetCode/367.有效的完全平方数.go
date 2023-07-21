/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num <= 4 {
		if num == 1 || num == 4 {
			return true
		} else {
			return false
		}
	}
	left := 0
	right := num / 2
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// @lc code=end

