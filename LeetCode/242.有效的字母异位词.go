/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	record := [26]int{}
	for _, r := range s {
		record[r-rune('a')]++
	}
	for _, r := range t {
		record[r-rune('a')]--
	}
	// record数组如果有的元素不为零0，说明字符串s和t 一定是谁多了字符或者谁少了字符。
	return record == [26]int{}
}

// @lc code=end

