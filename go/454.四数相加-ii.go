/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int) // key: a+b的值，value: a+b数值出现的次数
	count := 0
	// 遍历nums1和nums2，统计a+b的值以及出现的次数
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}
	// 遍历nums3和nums4，寻找map中(-c-d)的值，如果存在则count加上对应的次数
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += m[-v3-v4]
		}
	}
	return count
}

// @lc code=end

