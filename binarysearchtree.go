package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type BinarySearchTree[T cmp.Ordered] struct {
	head *node[T]
}

type node[T cmp.Ordered] struct {
	value T
	left  *node[T]
	right *node[T]
}

func (bst BinarySearchTree[T]) Insert(value T) {
	bst.head.insert(value)
}

func (n *node[T]) insert(value T) {
	if n == nil {
		n = &node[T]{value, nil, nil}
	} else if value < n.value {
		n.left.insert(value)
	} else {
		n.right.insert(value)
	}
}

func (bst BinarySearchTree[T]) inOrderTraversal() []T {
	return bst.head.inOrderTraversal()
}

func (n *node[T]) inOrderTraversal() []T {
	if n == nil {
		return []T{}
	} else {
		return slices.Concat(n.left.inOrderTraversal(), []T{n.value}, n.right.inOrderTraversal())
	}
}

func (bst BinarySearchTree[T]) String() string {
	return strings.Trim(
		strings.Join(
			strings.Fields(
				fmt.Sprint(bst.inOrderTraversal()),
			),
			" ",
		),
		"[]",
	)
}
