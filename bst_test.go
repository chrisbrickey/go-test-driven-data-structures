package goBST

import (
	"testing"
	"fmt"
)


func TestEmptyTreeHasNilRoot(t *testing.T) {
	var tree = BST{}
	emptyRoot := tree.root

	if emptyRoot != nil {
		t.Errorf("TestEmptyTreeHasNilRoot: tree.root failed")
		fmt.Println("Expected: ", nil)
		fmt.Println("Actual: ", emptyRoot)
	}

}
