package binarytree

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
