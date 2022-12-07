package main

import "math"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func minHead(lists []*ListNode) (min int, index int) {
	min = math.MaxInt
	index = -1

	for i, node := range lists {
		if node != nil && node.Val < min {
			min = node.Val
			index = i
		}
	}
	return min, index
}

func mergeKLists(lists []*ListNode) *ListNode {
	min, index := minHead(lists)
	if index == -1 {
		return nil
	}
	head := &ListNode{Val: min}
	lists[index] = lists[index].Next
	current := head

	min, index = minHead(lists)
	for index != -1 {
		current.Next = &ListNode{Val: min}
		lists[index] = lists[index].Next
		min, index = minHead(lists)
		current = current.Next
	}
	return head
}
