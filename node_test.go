package goBST

import (
	"testing"
	"fmt"
)

func TestNodeConstructorWithInput(t *testing.T) {
	var node = Node{1}
	valueResult := node.val

	if valueResult != 1 {
		t.Errorf("TestNodeConstructorWithInput: node.val failed")
		fmt.Println("Expected: ", 1)
		fmt.Println("Actual: ", valueResult)
	}

}