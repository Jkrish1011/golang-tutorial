package leetcode

import "fmt"

func strStr(haystack string, needle string) int {
	idx_haystack := 0
	var array_of_index []int

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			array_of_index = append(array_of_index, i)
		}
	}
	fmt.Printf("array_of_index: %v\n", array_of_index)

	for _, value := range array_of_index {
		// Checking if the evaluation will result in index-out-of-bound exception
		if value+len(needle) > len(haystack) {
			break
		}
		var foundFlag bool = true
		idx_haystack = value
		for _, char := range needle {
			if haystack[idx_haystack] == byte(char) {
				fmt.Printf("found\n")
				idx_haystack += 1
			} else {
				foundFlag = false
				idx_haystack = 0
				break
			}
		}

		if foundFlag {
			fmt.Printf("Found the needle: idx :: %v\n", value)
			return value
		}
	}

	return -1
}
