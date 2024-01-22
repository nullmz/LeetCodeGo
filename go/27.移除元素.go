/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start

// 前后指针法(前指针遇到val值时，从后面找一个不是val值的数值进行替换)
// func removeElement(nums []int, val int) int {
// 	pre, end := 0, len(nums)-1
// 	for pre <= end {
// 		for pre <= end && nums[pre] != val {
// 			pre++
// 		}
// 		for pre <= end && nums[end] == val {
// 			end--
// 		}
// 		if pre < end {
// 			nums[pre] = nums[end]
// 			pre++
// 			end--
// 		}
// 	}
// 	fmt.Println(nums)
// 	return pre
// }

// 快慢指针法(fastIndex替换slowIndex的内容，当fastIndex遇到val值时不执行替换操作)
func removeElement(nums []int, val int) int {
	slowIndex := 0
	lens := len(nums)
	for fastIndex := 0; fastIndex < lens; fastIndex++ {
		if nums[fastIndex] != val {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	fmt.Println(nums)
	return slowIndex
}

// @lc code=end

