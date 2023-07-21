/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if x <= 4 {

		if x == 4 {
			return 2
		} else if x == 0 {
			return 0
		} else {
			return 1
		}
	}
	min := 0
	max := x / 2
	for min <= max {
		mid := (max + min) / 2
		if mid*mid > x {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return max
}

// @lc code=end

