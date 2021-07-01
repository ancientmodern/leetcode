package recurse

import (
	. "LinkedList/util"
)

// ReverseList 反转整个链表
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

var succ *ListNode = nil

// ReverseListN 反转链表前 N 个元素
func ReverseListN(head *ListNode, n int) *ListNode {
	if n == 1 {
		succ = head.Next
		return head
	}
	last := ReverseListN(head.Next, n-1)
	head.Next.Next = head
	head.Next = succ
	return last
}

// ReverseBetween 反转链表 [left, right] 的元素
func ReverseBetween(head *ListNode, left, right int) *ListNode {
	if left == 1 {
		return ReverseListN(head, right)
	}
	head.Next = ReverseBetween(head.Next, left-1, right-1)
	return head
}
