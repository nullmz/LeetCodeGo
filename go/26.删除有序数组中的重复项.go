/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start

// 快慢指针（如果fast==fast-1，说明遇到重复数值，则fast继续往后找到不重复的值再更新nums[slow]的值
func removeDuplicates(nums []int) int {
	lens:=len(nums)
	if lens==0{
		return 0
	}
	slow:=1
	for fast:=1;fast<lens;fast++{
		if nums[fast]!=nums[fast-1]{
			nums[slow]=nums[fast]
			slow++
		}
	}
	return slow
}
// @lc code=end

