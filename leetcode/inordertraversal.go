package leetcode

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:
Input: root = [1,null,2,3]
Output: [1,3,2]


Example 2:
Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
Output: [4,2,6,5,7,1,3,9,8]


Example 3:
Input: root = []
Output: []

Example 4:
Input: root = [1]
Output: [1]

Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal_RecursiveApproach(root *TreeNode) []int {

	inOrderTraversed := []int{}
	var depthFirstSearch func(*TreeNode)
	depthFirstSearch = func(root *TreeNode) {
		if root == nil {
			return
		}
		depthFirstSearch(root.Left)
		inOrderTraversed = append(inOrderTraversed, root.Val)
		depthFirstSearch(root.Right)
	}
	depthFirstSearch(root)

	return inOrderTraversed
}

// Stack based approach
func inorderTraversal(root *TreeNode) []int {
	stack := NewStack[TreeNode]()

	inOrderTraversed := []int{}
	node := root

	for len(*stack) != 0 || node != nil {
		if node != nil {
			stack.Push(*node)
			node = node.Left
			continue
		}
		topNode := stack.Pop()
		inOrderTraversed = append(inOrderTraversed, topNode.Val)
		node = topNode.Right
	}

	return inOrderTraversed
}
