package tree

import (
	"fmt"
	"github.com/mats9693/utils/support"
	"sort"
	"testing"
)

func TestNewBST(t *testing.T) {
	data := support.GenerateRandomIntSlice(20, 100)

	bstIns := newBST(data)

	if !isBST(bstIns) {
		t.Logf("data: %v\ntree: \n%s", data, printBST(bstIns))
		t.Fail()
	}
}

func TestBinarySearchTree_Find(t *testing.T) {
	data := support.GenerateRandomIntSlice(20, 100, 0, 10, 101)

	bstIns := newBST(data)

	values := []int{0, 10, 101, 200}
	expected := []bool{true, true, true, false}

	if len(values) != len(expected) {
		t.Logf("unmatched amount of values and expected results")
		t.Fail()
	}

	for i := range values {
		res := bstIns.Find(values[i])
		if res != expected[i] {
			t.Logf("index: %d, find: %d\n\ttree: \n%s\n\twant\n\t%T\n\tget: %T\n",
				i, values[i], printBST(bstIns), expected[i], res)
			t.Fail()
		}
	}
}

func printBST(tree *binarySearchTreeImpl) string {
	if tree == nil {
		return ""
	}

	nodeList := []*bstNode{tree.root}
	res := ""
	for len(nodeList) > 0 {
		nextNodeList := make([]*bstNode, 0, len(nodeList)*2)

		for len(nodeList) > 0 {
			node := nodeList[0]
			nodeList = nodeList[1:]

			if node == nil {
				res += "null "
			} else {
				res += fmt.Sprintf("%d ", node.value)
				nextNodeList = append(nextNodeList, node.left, node.right)
			}
		}

		res += "\n"
		nodeList = nextNodeList
	}

	return res
}

func isBST(tree *binarySearchTreeImpl) bool {
	if tree == nil {
		return true
	}

	values := dfs(tree.root)
	backup := support.DeepCopyIntSlice(values)

	sort.Ints(values)

	return support.CompareOnIntSlice(backup, values)
}

func dfs(root *bstNode) []int {
	if root == nil {
		return nil
	} else if root.left == nil && root.right == nil {
		return []int{root.value}
	}

	leftValues := dfs(root.left)

	rightValues := dfs(root.right)

	res := make([]int, 0, len(leftValues)+len(rightValues)+1)

	res = append(res, leftValues...)
	res = append(res, root.value)
	res = append(res, rightValues...)

	return res
}
