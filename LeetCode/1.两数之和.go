/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// 暴力解法
// func twoSum(nums []int, target int) []int {
//     for k1, _ := range nums {
//         for k2 := k1 + 1; k2 < len(nums); k2++ {
//             if target == nums[k1] + nums[k2] {
//                 return []int{k1, k2}
//             }
//         }
//     }
//     return []int{}
// }

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}
	return []int{}
}

// @lc code=end

