package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	var nums1 []int = []int{1, 2, 3, 0, 0, 0}
	var nums2 []int = []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Printf("%+v\n", nums1)
}
