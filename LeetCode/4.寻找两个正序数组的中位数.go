/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m == 0 {
		if n%2 == 0 {
			return (float64(nums2[n/2-1]+nums2[n/2]) / 2.0)
		} else {
			return float64(nums2[n/2])
		}
	} else if n == 0 {
		if m%2 == 0 {
			return (float64(nums1[m/2-1]+nums1[m/2]) / 2.0)
		} else {
			return float64(nums1[m/2])
		}
	}

	num := make([]int, m+n)
	// var num [m + n]float64
	count := 0
	i, j := 0, 0
	for count != (m + n) {
		if i == m {
			for j != n {
				num[count] = nums2[j]
				count++
				j++
			}
			break
		}
		if j == n {
			for i != m {
				num[count] = nums1[i]
				count++
				i++
			}
			break
		}

		if nums1[i] < nums2[j] {
			num[count] = nums1[i]
			count++
			i++
		} else {
			num[count] = nums2[j]
			count++
			j++
		}
	}
	if count%2 == 0 {
		return (float64(num[count/2-1]+num[count/2]) / 2.0)
	} else {
		return float64(num[count/2])
	}
}

// @lc code=end

