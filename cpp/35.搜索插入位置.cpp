/*
 * @lc app=leetcode.cn id=35 lang=cpp
 *
 * [35] 搜索插入位置
 */

// @lc code=start
// 暴力解法(遍历)
// class Solution
// {
// public:
//     int searchInsert(vector<int> &nums, int target)
//     {
//         for (int i = 0; i < nums.size(); i++)
//         {
//             // 如果当前元素大于等于目标值，返回当前下标
//             if (nums[i] >= target)
//             {
//                 return i;
//             }
//         }
//         return nums.size(); // 如果遍历完数组都没有找到大于等于目标值的元素，返回数组长度
//     }
// };

// 二分法
class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        int n = nums.size();
        int left = 0;
        int right = n - 1; // 定义target在左闭右闭的区间里，[left, right]
        while (left <= right) { 
            int middle = left + ((right - left) / 2);
            if (nums[middle] > target) {
                right = middle - 1; 
            } else if (nums[middle] < target) {
                left = middle + 1; 
            } else { 
                return middle;
            }
        }
        return right + 1;
    }
};
// @lc code=end
