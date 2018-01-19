package goBST

import (
	"testing"
	"fmt"
)

func TestNodeConstructorWithInput(t *testing.T) {
	var node = Node{1, nil, nil, nil}
	valueResult := node.value

	if valueResult != 1 {
		t.Errorf("TestNodeConstructorWithInput: node.value failed")
		fmt.Println("Expected: ", 1)
		fmt.Println("Actual: ", valueResult)
	}

}

func TestNodeConstructorChildrenParentNil(t *testing.T) {
	var node = Node{1, nil, nil, nil}
	rightResult := node.right
	leftResult := node.left
	parentResult := node.parent

	if rightResult != nil {
		t.Errorf("TestNodeConstructorRightNil: node.right failed")
		fmt.Println("Expected: ", nil)
		fmt.Println("Actual: ", rightResult)
	}

	if leftResult != nil {
		t.Errorf("TestNodeConstructorLeftNil: node.left failed")
		fmt.Println("Expected: ", nil)
		fmt.Println("Actual: ", leftResult)
	}

	if parentResult != nil {
		t.Errorf("TestNodeConstructorParentNil: node.parent failed")
		fmt.Println("Expected: ", nil)
		fmt.Println("Actual: ", parentResult)
	}


}