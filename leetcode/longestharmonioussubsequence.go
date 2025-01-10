package leetcode

/*
We define a harmonious array as an array where the difference between its maximum value and its minimum value is exactly 1.

Given an integer array nums, return the length of its longest harmonious
subsequence
 among all its possible subsequences.



Example 1:

Input: nums = [1,3,2,2,5,2,3,7]

Output: 5

Explanation:

The longest harmonious subsequence is [3,2,2,2,3].

Example 2:

Input: nums = [1,2,3,4]

Output: 2

Explanation:

The longest harmonious subsequences are [1,2], [2,3], and [3,4], all of which have a length of 2.

Example 3:

Input: nums = [1,1,1,1]

Output: 0

Explanation:

No harmonic subsequence exists.



Constraints:

1 <= nums.length <= 2 * 104
-109 <= nums[i] <= 109

*/

import "math"

func findLHS(nums []int) int {
	lookUp := make(map[int]int)
	maxValue := 0
	for _, num := range nums {
		_, ok := lookUp[num]
		if ok {
			lookUp[num] = lookUp[num] + 1
		} else {
			lookUp[num] = 1
		}
	}
	// fmt.Printf("%+v\n", lookUp)

	for num, val := range lookUp {
		nextVal, ok := lookUp[num+1]
		if ok {
			maxValue = int(math.Max(float64(maxValue), float64(val+nextVal)))
		}
	}
	return maxValue
}
