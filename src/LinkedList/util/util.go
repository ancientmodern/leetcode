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

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}
