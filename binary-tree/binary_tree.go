package binarytree

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (t *BinaryTree) Insert(val int) {
	t.Root = insert(t.Root, val)
}

func insert(current *Node, val int) *Node {
	if current == nil {
		return &Node{Val: val}
	}

	if val < current.Val {
		current.Left = insert(current.Left, val)
	} else if val > current.Val {
		current.Right = insert(current.Right, val)
	}

	// if val == current.Val Ignore duplicate values

	return current
}

func (t *BinaryTree) Print() {
	if t.Root == nil {
		return
	}

	queue := []*Node{t.Root}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := range levelSize {
			current := queue[0]
			queue = queue[1:]

			if current.Left != nil {
				queue = append(queue, current.Left)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}

			if i > 0 {
				fmt.Print(" - ")
			}

			fmt.Print(current.Val)
		}

		fmt.Println()
	}

}
