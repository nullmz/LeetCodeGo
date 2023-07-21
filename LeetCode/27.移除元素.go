/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	pre, end := 0, len(nums)-1
	for pre <= end {
		for pre <= end && nums[pre] != val {
			pre++
		}
		for pre <= end && nums[end] == val {
			end--
		}
		if pre < end {
			nums[pre] = nums[end]
			pre++
			end--
		}
	}
	fmt.Println(nums)
	return pre
}

// @lc code=end

