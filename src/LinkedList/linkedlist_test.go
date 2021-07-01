package main

import (
	"LinkedList/iteration"
	"LinkedList/recurse"
	. "LinkedList/util"
	"testing"
)

func equal(actual, expected *ListNode) bool {
	for actual != nil && expected != nil {
		if actual.Val != expected.Val {
			return false
		}
		actual = actual.Next
		expected = expected.Next
	}
	if (actual == nil && expected != nil) || (actual != nil && expected == nil) {
		return false
	}
	return true
}

func TestRecurse(t *testing.T) {
	var (
		in    = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		left  = 4
		right = 8
		out   = []int{1, 2, 3, 8, 7, 6, 5, 4, 9, 10, 11, 12}
	)
	head := CreateList(in)
	actual := recurse.ReverseBetween(head, left, right)
	expected := CreateList(out)
	if !equal(actual, expected) {
		t.Errorf("recurse.ReverseBetween(%v, %d, %d) = %v; expected %v", in, left, right, List2Slice(actual), out)
	}
}

func TestIteration(t *testing.T) {
	var (
		in  = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		k   = 3
		out = []int{3, 2, 1, 6, 5, 4, 9, 8, 7, 12, 11, 10}
	)
	head := CreateList(in)
	actual := iteration.ReverseKGroup(head, k)
	expected := CreateList(out)
	if !equal(actual, expected) {
		t.Errorf("iteration.ReverseKGroup(%v, %d) = %v; expected %v", in, k, List2Slice(actual), out)
	}
}
