/*
 * @lc app=leetcode.cn id=904 lang=golang
 *
 * [904] 水果成篮
 */

// @lc code=start
func totalFruit(fruits []int) int {
	cnt := map[int]int{}
	pre := 0
	result := 0
	for end, x := range fruits {
		cnt[x]++
		for len(cnt) > 2 {
			y := fruits[pre]
			cnt[y]--
			if cnt[y] == 0 {
				delete(cnt, y)
			}
			pre++
		}
		result = max(result, end-pre+1)
	}
	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

