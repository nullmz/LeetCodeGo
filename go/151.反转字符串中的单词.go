/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
// 方法一
func reverseWords(s string) string {
	b := []byte(s)
	// 移除前面、中间、后面的多余空格
	slow := 0
	for i := 0; i < len(b); i++ {
		if b[i] != ' ' {
			if slow != 0 {
				b[slow] = ' '
				slow++
			}
			for i < len(b) && b[i] != ' ' { //复制逻辑
				b[slow] = b[i]
				slow++
				i++
			}
		}
	}

	b = b[0:slow]

	// 翻转整个字符串
	reverse(b)
	// 翻转每个单词
	last := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b[last:i])
			last = i + 1
		}
	}
	return string(b)
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

// 方法二
// func reverseWords(s string) string {
// 	//1.使用双指针删除冗余的空格
// 	slowIndex, fastIndex := 0, 0
// 	b := []byte(s)
// 	//删除头部冗余空格
// 	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
// 		fastIndex++
// 	}
// 	//删除单词间冗余空格
// 	for ; fastIndex < len(b); fastIndex++ {
// 		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
// 			continue
// 		}
// 		b[slowIndex] = b[fastIndex]
// 		slowIndex++
// 	}
// 	//删除尾部冗余空格
// 	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
// 		b = b[:slowIndex-1]
// 	} else {
// 		b = b[:slowIndex]
// 	}
// 	//2.反转整个字符串
// 	reverse(&b, 0, len(b)-1)
// 	//3.反转单个单词  i单词开始位置，j单词结束位置
// 	i := 0
// 	for i < len(b) {
// 		j := i
// 		for ; j < len(b) && b[j] != ' '; j++ {
// 		}
// 		reverse(&b, i, j-1)
// 		i = j
// 		i++
// 	}
// 	return string(b)
// }

// func reverse(b *[]byte, left, right int) {
// 	for left < right {
// 		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
// 		left++
// 		right--
// 	}
// }

// @lc code=end

