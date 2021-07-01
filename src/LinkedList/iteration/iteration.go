package iteration

import (
	. "LinkedList/util"
)

// ReverseList 反转整个链表
func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur, next := head, head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// ReverseBetween 反转链表 [head, end) 的元素
func ReverseBetween(head, end *ListNode) *ListNode {
	var pre *ListNode
	cur, next := head, head
	for cur != end {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// ReverseKGroup K个一组反转链表
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	start, end := head, head
	for i := 0; i < k; i++ {
		// 剩余不足k个，不做操作直接返回
		if end == nil {
			return head
		}
		end = end.Next
	}
	newHead := ReverseBetween(start, end)
	start.Next = ReverseKGroup(end, k)
	return newHead
}
