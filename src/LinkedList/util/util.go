package util

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(array []int) *ListNode {
	head := new(ListNode)
	p := head
	for i, v := range array {
		if i == 0 {
			head.Val = v
		} else {
			p.Next = new(ListNode)
			p = p.Next
			p.Val = v
		}
	}
	p = nil
	return head
}

func List2Slice(head *ListNode) []int {
	ans := make([]int, 0)
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	return ans
}

func PrintList(head *ListNode) {
	fmt.Println(List2Slice(head))
}
