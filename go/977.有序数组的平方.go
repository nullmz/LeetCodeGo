/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start

// 方法一：nums中的数平方后直接排序
// func sortedSquares(nums []int) []int {
// 	ans := make([]int, len(nums))
// 	for i, v := range nums {
// 		ans[i] = v * v
// 	}
// 	sort.Ints(ans)
// 	return ans
// }

// 方法二：双指针（i左边一半为负数，j右边一半为非负数，双指针计算小的平方数放入ans切片）
// func sortedSquares(nums []int) []int {
// 	n := len(nums)
// 	lastNegIndex := -1
// 	for i := 0; i < n && nums[i] < 0; i++ {
// 		lastNegIndex = i
// 	}
// 	ans := make([]int, 0, n)
// 	for i, j := lastNegIndex, lastNegIndex+1; i >= 0 || j < n; {
// 		if i < 0 {
// 			ans = append(ans, nums[j]*nums[j])
// 			j++
// 		} else if j == n {
// 			ans = append(ans, nums[i]*nums[i])
// 			i--
// 		} else if nums[i]*nums[i] < nums[j]*nums[j] {
// 			ans = append(ans, nums[i]*nums[i])
// 			i--
// 		} else {
// 			ans = append(ans, nums[j]*nums[j])
// 			j++
// 		}
// 	}
// 	return ans
// }

// 方法三：双指针，将答案逆序放入，不用考虑指针移动到边界的情况
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	pre, end := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		if v, w := nums[pre]*nums[pre], nums[end]*nums[end]; v > w {
			ans[pos] = v
			pre++
		} else {
			ans[pos] = w
			end--
		}
	}
	return ans
}

// @lc code=end

