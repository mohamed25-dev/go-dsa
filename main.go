package main

import (
	binarytree "github.com/mohamed25-dev/go-dsa/binary-tree"
)

func main() {

	var t binarytree.BinaryTree
	nums := []int{9, 7, 5, 11, 10, 12, 8}

	for _, v := range nums {
		t.Insert(v)
	}

	t.Print()
}
