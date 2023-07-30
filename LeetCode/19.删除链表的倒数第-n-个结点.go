/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := head
	pre := dummyHead
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			pre = pre.Next
		}
		i++
	}
	pre.Next = pre.Next.Next
	return dummyHead.Next
}

// @lc code=end

