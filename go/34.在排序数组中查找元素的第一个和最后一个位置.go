/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
/*
情况一：target在数组范围的左边之外或右边之外
情况二：target在数组范围内但是数组中不存在target
情况三：target在数组中
*/
// 返回-1表示出错
func getLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// situation 1
	// right==-1要考虑数组为空时候的情况
	if right == -1 || target < nums[0] || target > nums[right] {
		return -1
	}
	for left <= right {
		// mid := left + (right-left)/2
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if nums[left] == target {
		// situation 3
		return left
	} else {
		// situation 2
		return -1
	}
}
func getRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// situatuion one
	if target < nums[0] || target > nums[right] {
		return -1
	}
	for left <= right {
		// mid := left + (right-left)/2
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if nums[right] == target {
		// situation three
		return right
	} else {
		// situation two
		return -1
	}
}

func searchRange(nums []int, target int) []int {
	leftBorder := getLeft(nums, target)
	if leftBorder == -1 {
		return []int{-1, -1}
	} else {
		rightBorder := getRight(nums, target)
		return []int{leftBorder, rightBorder}
	}

}

// @lc code=end

