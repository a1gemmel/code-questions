package main

import (
	"fmt"
	"math"
)

type bst struct {
	val   int
	count int
	left  *bst
	right *bst
}

func (b *bst) insert(val int) *bst {
	if b == nil {
		return &bst{
			val:   val,
			count: 1,
		}
	}
	if b.val == val {
		b.count++
		return b
	}
	if val < b.val {
		b.left = b.left.insert(val)
	} else {
		b.right = b.right.insert(val)
	}
	return b
}

func (b *bst) min() int {
	if b == nil {
		return int(math.Inf(1))
	}
	if b.left == nil {
		return b.val
	}
	return b.left.min()
}

func (b *bst) max() int {
	if b == nil {
		return int(math.Inf(-1))
	}
	if b.right == nil {
		return b.val
	}
	return b.right.max()
}

func (b *bst) remove(val int) *bst {
	if b == nil {
		return nil
	}
	if b.val == val {
		if b.count > 1 {
			b.count--
			return b
		} else {
			if b.right == nil {
				return b.left
			}
			if b.left == nil {
				return b.right
			}

			leftMostChild := b.right

			for leftMostChild.left != nil {
				leftMostChild = leftMostChild.left
			}
			leftMostChild.left = b.left
			return b.right
		}
	}
	if val < b.val {
		b.left = b.left.remove(val)
	} else {
		b.right = b.right.remove(val)
	}
	return b
}

func (b *bst) String() string {
	if b == nil {
		return "[]"
	}
	return fmt.Sprintf("[ %d(%d) %s , %s]", b.val, b.count, b.left.String(), b.right.String())
}

func countSegments(arr []int, threshold int) int {
	var tree *bst = nil

	start := 0
	end := 0
	// This algorithm won't count the last single length segment, so start at 1
	segments := 1
	tree = tree.insert(arr[0])

	for start < len(arr) {
		if tree.max()-tree.min() <= threshold && end+1 < len(arr) {
			end++
			tree = tree.insert(arr[end])
		} else {
			// We've found all segments that start at 'start'
			fmt.Printf("Max segment length at %d is %d (start %d end %d)\n", start, end-start, start, end-1)
			segments += end - start
			tree = tree.remove(arr[start])
			start++
			if end+1 < len(arr) {
				end++
			}
		}
	}

	return segments
}
