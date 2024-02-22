package main

import (
	"fmt"
)

func main() {
	var bst BinarySearchTree[int]

	bst.Insert(12)
	bst.Insert(7)
	bst.Insert(15)

	fmt.Println(bst)
}
