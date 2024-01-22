/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 双指针(走完最长串长度的一趟，差值为两串长的之差)
// pointA从headA链跳转到pointB链，
// pointB从headB链跳到pointB链，
// 跳转之后的双指针之后的值相当于右对齐
// 如果不相交，pointA和pointB到链表末尾时都为nil，退出循环
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pointA, pointB := headA, headB
	for pointA != pointB {
		if pointA != nil {
			pointA = pointA.Next
		} else {
			pointA = headB
		}

		if pointB != nil {
			pointB = pointB.Next
		} else {
			pointB = headA
		}
	}
	return pointA
}

// 普通方法，先求长度再出发比较
// fast指向长链表，slow指向短链表
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	curA, curB := headA, headB
// 	lenA, lenB := 0, 0
// 	// 求A,B长度
// 	for curA != nil {
// 		curA = curA.Next
// 		lenA++
// 	}
// 	for curB != nil {
// 		curB = curB.Next
// 		lenB++
// 	}

// 	var step int
// 	var fast, slow *ListNode
// 	// 求出长度差，并且让更长的链表先走相差的长度
// 	if lenA > lenB {
// 		step = lenA - lenB
// 		fast, slow = headA, headB
// 	} else {
// 		step = lenB - lenA
// 		fast, slow = headB, headA
// 	}

// 	for i := 0; i < step; i++ {
// 		fast = fast.Next
// 	}
// 	// 遍历两链表，遇到相同的退出
// 	for fast != slow {
// 		fast = fast.Next
// 		slow = slow.Next
// 	}
// 	return fast
// }

// @lc code=end

