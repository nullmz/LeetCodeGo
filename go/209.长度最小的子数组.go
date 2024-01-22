/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	pre := 0
	lens := len(nums)
	sum := 0
	result := lens + 1
	for end := 0; end < lens; end++ {
		sum += nums[end]
		for sum >= target {
			subLens := end - pre + 1
			if subLens < result {
				result = subLens
			}
			sum -= nums[pre]
			pre++
		}
	}
	if result == lens+1 {
		return 0
	} else {
		return result
	}
}

// @lc code=end

