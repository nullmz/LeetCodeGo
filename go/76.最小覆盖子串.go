/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	// 定义两个map
	ori, cnt := map[byte]int{}, map[byte]int{}
	// 计算t中的字母
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	lens := len(s)
	// 获取最大的整数值
	result := math.MaxInt32
	resPre, resEnd := -1, -1

	check := func() bool {
		for key, val := range ori {
			if cnt[key] < val {
				return false
			}
		}
		return true
	}

	for pre, end := 0, 0; end < lens; end++ {
		if end < lens && ori[s[end]] > 0 {
			cnt[s[end]]++
		}
		for check() && pre <= end {
			if (end - pre + 1) < result {
				result = end - pre + 1
				resPre, resEnd = pre, pre + result
			}
			if _, ok := ori[s[pre]]; ok {
				cnt[s[pre]] -= 1
			}
			pre++
		}
	}
	if resPre == -1 {
		return ""
	}
	return s[resPre:resEnd]
}

// @lc code=end

