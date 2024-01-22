/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{}, 0) //用map模拟set
	res := make([]int, 0)
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {	//判断集合 set 中是否不存在键 v
			set[v] = struct{}{}
		}
	}

	for _, v := range nums2 {
		// 如果存在与一个数组中，则加入结果集，并清空该set
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}

// @lc code=end

