package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddTwoNumbers(t *testing.T) {
	tests := []struct {
		a        *ListNode
		b        *ListNode
		expected *ListNode
	}{
		{
			a:        &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			b:        &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			expected: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			a:        &ListNode{Val: 9},
			b:        &ListNode{Val: 9, Next: &ListNode{Val: 9}},
			expected: &ListNode{Val: 8, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
	}

	for _, test := range tests {
		assert.True(t, test.expected.Equals(addTwoNumbers(test.a, test.b)))
	}

}
