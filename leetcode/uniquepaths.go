package leetcode

/*
There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]).
The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.

Example 1:
Input: m = 3, n = 7
Output: 28

Example 2:
Input: m = 3, n = 2
Output: 3

Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down


Constraints:

1 <= m, n <= 100
*/

func uniquePaths(m int, n int) int {
	row := make([]int, n, n)
	for i := range row {
		row[i] = 1
	}
	newRow := row
	// fmt.Printf("%+v\n", newRow)
	// fmt.Printf("%+v\n", row)
	for rowIdx := m - 2; rowIdx >= 0; rowIdx-- {
		for colIdx := n - 2; colIdx >= 0; colIdx-- {
			newRow[colIdx] = newRow[colIdx+1] + row[colIdx]
		}
		row = newRow
	}
	return row[0]
}
