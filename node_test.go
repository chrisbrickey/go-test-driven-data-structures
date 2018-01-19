package goBST

import (
	"testing"
	"fmt"
)

func TestNodeConstructorWithoutInput(t *testing.T) {
	var node = Node{}
	valueResult := node.val

	if valueResult != nil {
		t.Errorf("TestNodeConstructorWithoutInput: node.val failed")
		fmt.Println("Expected: ", nil)
		fmt.Println("Actual: ", valueResult)
	}

}