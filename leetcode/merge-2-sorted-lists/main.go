package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		head = &ListNode{Val: list1.Val}
		list1 = list1.Next
	} else {
		head = &ListNode{Val: list2.Val}
		list2 = list2.Next
	}

	current := head

	for list1 != nil || list2 != nil {
		if list1 == nil {
			current.Next = list2
			break
		}
		if list2 == nil {
			current.Next = list1
			break
		}
		if list1.Val < list2.Val {
			current.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			current.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		current = current.Next
	}
	return head
}

func mergeTwoListsRecursion(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		return &ListNode{Val: list1.Val, Next: mergeTwoListsRecursion(list1.Next, list2)}
	}
	return &ListNode{Val: list2.Val, Next: mergeTwoListsRecursion(list1, list2.Next)}
}

func makeList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	return &ListNode{Val: arr[0], Next: makeList(arr[1:])}
}

func printList(l *ListNode) string {
	s := "[ "
	for l != nil {
		s += fmt.Sprintf("%d, ", l.Val)
		l = l.Next
	}
	return s + " ]"
}

func main() {
	fmt.Println(mergeTwoLists(makeList([]int{1, 2, 4}), makeList([]int{1, 3, 4})))
}
