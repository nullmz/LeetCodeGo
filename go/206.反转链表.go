/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// double pointer
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	cur := head.Next
// 	pre := head
// 	pre.Next = nil
// 	for cur != nil {
// 		next := cur.Next
// 		cur.Next = pre
// 		pre = cur
// 		cur = next
// 	}
// 	return pre
// }

// 双指针
func reverseList(head *ListNode) *ListNode {
	// pre.Next start with empty.(nice!)
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 递归(每次递归改变cur.Next一个节点的指向)
// func reverseList(head *ListNode) *ListNode {
// 	return help(nil, head)
// }
// func help(pre, head *ListNode) *ListNode {
// 	if head == nil {
// 		return pre
// 	}
// 	next := head.Next
// 	head.Next = pre
// 	return help(head, next)
// }

// @lc code=end

