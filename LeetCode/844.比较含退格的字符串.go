/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
// 栈，普通字符压栈，遇到#栈顶弹出
// func build(str string) string {
// 	s := []byte{}
// 	for i := range str {
// 		if str[i] != '#' {
// 			s = append(s, str[i])
// 		} else if len(s) > 0 {
// 			s = s[:len(s)-1]
// 		}
// 	}
// 	return string(s)
// }

// func backspaceCompare(s string, t string) bool {
// 	return build(s) == build(t)
// }

// 双指针
// 逆序遍历，如果每一个不被删除的元素都相等则true，否则false

func backspaceCompare(s string, t string) bool {
	skipS, skipT := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		// 跳过s[]中被删除的元素
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		// 跳过t[]中要被删除的元素
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}

		// 判断每一个不被删除的元素是否相同（倒序判断）
		if i >= 0 && j >= 0 {
			if s[i] != t[j] { //找到不同元素
				return false
			}
		} else if i >= 0 || j >= 0 { //其中一组已经结束
			return false
		}
		i--
		j--
	}
	return true
}

// @lc code=end

