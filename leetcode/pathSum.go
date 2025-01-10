package leetcode

/*
Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

A leaf is a node with no children.

Example 1:

Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
Output: true
Explanation: The root-to-leaf path with the target sum is shown.

Example 2:

Input: root = [1,2,3], targetSum = 5
Output: false
Explanation: There are two root-to-leaf paths in the tree:
(1 --> 2): The sum is 3.
(1 --> 3): The sum is 4.
There is no root-to-leaf path with sum = 5.

Example 3:

Input: root = [], targetSum = 0
Output: false
Explanation: Since the tree is empty, there are no root-to-leaf paths.


Constraints:

The number of nodes in the tree is in the range [0, 5000].
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "fmt"

func dfs(root *TreeNode, targetSum int, currSum int) bool {
	if root == nil {
		return false
	}

	fmt.Printf("root.Val+currSum :: %v\n", root.Val+currSum)
	fmt.Printf("targetSum :: %v\n", targetSum)
	fmt.Printf("Checking if Target Sum is a match \n")

	if root.Val+currSum == targetSum && root.Left == nil && root.Right == nil {
		fmt.Printf("Target Found!\n")
		return true
	}

	fmt.Printf("Taking Left\n")
	lf := dfs(root.Left, targetSum, currSum+root.Val)
	if lf == true {
		return true
	}
	fmt.Printf("Left tree unsuccessful\n")
	fmt.Printf("Taking Right\n")
	rh := dfs(root.Right, targetSum, currSum+root.Val)
	if rh == true {
		return true
	}
	fmt.Printf("Right tree unsuccessful\n")
	return false
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	fmt.Printf("Taking Main Left Tree\n")
	lf := dfs(root.Left, targetSum, root.Val)
	if lf == true {
		return true
	}
	fmt.Printf("Main Left tree unsuccessful\n")
	fmt.Printf("Taking Main Right Tree\n")
	rh := dfs(root.Right, targetSum, root.Val)
	if rh == true {
		return true
	}
	fmt.Printf("Main Right tree unsuccessful\n")
	return false
}
