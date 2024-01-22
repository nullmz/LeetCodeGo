/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	//head=list[i]
	//pre=list[i-1]
	pre := dummy
	for head != nil && head.Next != nil {
		pre.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		//pre=list[(i+2)-1]
		pre = head
		//head=list[(i+2)]
		head = next
	}
	return dummy.Next
}

// 递归
// func swapPairs(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil {
//         return head
//     }
//     next := head.Next
//     head.Next = swapPairs(next.Next)
//     next.Next = head
//     return next
// }
// @lc code=end

