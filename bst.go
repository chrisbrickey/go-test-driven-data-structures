package goBST

type BST struct {
	root *Node
}

func(b *BST) add(n int32) {
	var newNode = Node{n, nil, nil, nil}
	if b.root == nil {
		b.root = &newNode
		return
	}

	current := b.root //Node

	if n < current.value {
		if current.left != nil {
			current.left = &newNode
			return
		}
	}


	if current.right != nil {
		current.right = &newNode
		return
	}



	//for current != nil {
	//	if n > current.value {
	//
	//	}
	//}


}



