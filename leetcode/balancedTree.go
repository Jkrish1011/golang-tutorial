package leetcode

/*
Balanced Binary Tree

Given a binary tree, determine if it is height-balanced.
((A height-balanced binary tree is a binary tree in which the depth of the two
subtrees of every node never differs by more than one.))

Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: true
Example 2:


Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
Example 3:

Input: root = []
Output: true


Constraints:

The number of nodes in the tree is in the range [0, 5000].
-104 <= Node.val <= 104

*/

import "math"

// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func checkIfBalanced(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	heightLeft, balancedLeft := checkIfBalanced(node.Left)
	heightRight, balancedRight := checkIfBalanced(node.Right)

	var balanced bool = balancedLeft && balancedRight && (math.Abs(float64(heightRight-heightLeft)) <= 1)

	return 1 + int(math.Max(float64(heightLeft), float64(heightRight))), balanced

}

func isBalanced(root *TreeNode) bool {
	_, balanced := checkIfBalanced(root)
	return balanced
}
