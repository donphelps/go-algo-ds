package tree

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	value int
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(v int) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{value: v, left: nil, right: nil}
	} else {
		t.root.insert(v)
	}
	return t
}

func (n *BinaryNode) insert(v int) {

}
