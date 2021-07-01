package main

import (
	"LinkedList/iteration"
	. "LinkedList/util"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	head := CreateList(array)
	PrintList(head)
	head = iteration.ReverseKGroup(head, 3)
	PrintList(head)
}
