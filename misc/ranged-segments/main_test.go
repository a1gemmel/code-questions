package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BST(t *testing.T) {

	var tree *bst = nil

	for _, v := range []int{4, 1, 5, 7, 6, 3} {
		tree = tree.insert(v)
	}

	assert.Equal(t, "[ 4(1) [ 1(1) [] , [ 3(1) [] , []]] , [ 5(1) [] , [ 7(1) [ 6(1) [] , []] , []]]]", tree.String())
	tree = tree.insert(7)
	assert.Equal(t, "[ 4(1) [ 1(1) [] , [ 3(1) [] , []]] , [ 5(1) [] , [ 7(2) [ 6(1) [] , []] , []]]]", tree.String())

	assert.Equal(t, 1, tree.min())
	assert.Equal(t, 7, tree.max())

	tree = tree.remove(7)
	assert.Equal(t, "[ 4(1) [ 1(1) [] , [ 3(1) [] , []]] , [ 5(1) [] , [ 7(1) [ 6(1) [] , []] , []]]]", tree.String())
	tree = tree.remove(7)
	assert.Equal(t, "[ 4(1) [ 1(1) [] , [ 3(1) [] , []]] , [ 5(1) [] , [ 6(1) [] , []]]]", tree.String())

	tree = nil

	for _, v := range []int{2, 1, 3} {
		tree = tree.insert(v)
	}
	tree = tree.remove(2)
	assert.Equal(t, "[ 3(1) [ 1(1) [] , []] , []]", tree.String())

}

func Test_CountSegments(t *testing.T) {
	cases := []struct {
		array     []int
		threshold int
		result    int
	}{
		{
			array:     []int{5, 2, 4, 3, 1, 6},
			threshold: 3,
			result:    15,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.result, countSegments(c.array, c.threshold))
	}
}
