package goBST

import (
	"testing"
	"fmt"
)

func Setup() BST {
	var tree = BST{}
	tree.add(11)
	return tree
}

func TestEmptyTreeHasNilRoot(t *testing.T) {
	var tree = BST{}
	actual := tree.root

	if actual != nil {
		t.Errorf("TestEmptyTreeHasNilRoot: tree.root failed")
		fmt.Println("Expected: ", nil)
		fmt.Println("Actual: ", actual)
	}

}


func TestRootNotNilAfterAdd(t *testing.T) {
	var tree = Setup()
	actual := tree.root

	if actual == nil {
		t.Errorf("TestRootNotNilAfterAdd: tree.root failed")
		fmt.Println("Expected: ", nil)
		fmt.Println("Actual: ", actual)
	}

}

func TestFirstValueAssignedToRoot(t *testing.T) {
	var tree = Setup()
	actual := tree.root.value
	var expectation int32 = 11

	if actual != expectation {
		t.Errorf("TestFirstValueAssignedToRoot: tree.root.value failed")
		fmt.Println("Expected: ", expectation)
		fmt.Println("Actual: ", actual)
	}

}

func TestSmallerValueInsertedToLeft(t *testing.T) {
	var tree = Setup()
	tree.add(10)
	var leftNode = tree.root.left

	if leftNode == nil {

	}
	actual := tree.root.left.value //left.value NOT valid if .left is nil
	var expectation int32 = 10

	if actual != expectation {
		t.Errorf("TestSmallerValueInsertedToLeft: tree.root.left.value failed")
		fmt.Println("Expected: ", expectation)
		fmt.Println("Actual: ", actual)
	}
}

func TestLargerValueInsertedToRight(t *testing.T) {
	var tree = Setup()
	tree.add(12)
	actual := tree.root.right.value
	var expectation int32 = 12

	if actual != expectation {
		t.Errorf("TestLargerValueInsertedToRight: tree.root.right.value failed")
		fmt.Println("Expected: ", expectation)
		fmt.Println("Actual: ", actual)
	}
}

func TestEquivalentValueInsertedToRight(t *testing.T) {
	var tree = Setup()
	tree.add(11)
	actual := tree.root.right.value
	var expectation int32 = 11

	if actual != expectation {
		t.Errorf("TestEquivalentValueInsertedToRight: tree.root.right.value failed")
		fmt.Println("Expected: ", expectation)
		fmt.Println("Actual: ", actual)
	}
}


