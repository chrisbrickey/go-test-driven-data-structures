package goBST

type BST struct {
	root *Node
}

func(b *BST) add(n int32) {
	var newNode = Node{n, nil, nil, nil}
	if b.root == nil {
		b.root = &newNode
	}
}



