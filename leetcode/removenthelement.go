package leetcode

/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:
----------
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
----------
Input: head = [1], n = 1
Output: []

Example 3:
----------
Input: head = [1,2], n = 1
Output: [1]

*/

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func removeNthFromEnd(head *ListNode1, n int) *ListNode1 {
	currNode := head
	prevNode := &ListNode1{Val: -1, Next: nil}
	// retVal := &ListNode1{Val: -1, Next: nil}
	// nextNode := currNode.Next
	counter := 0
	/*
		define a function to recursively iterate through the linkedlist using 3 pointers.
		Once we reach the end, we count back the n steps required.
		do prev->next = curr->next
		break
	*/
	var dp func(prevNode *ListNode1, currNode *ListNode1) (*ListNode1, *ListNode1)
	dp = func(prevNode *ListNode1, currNode *ListNode1) (*ListNode1, *ListNode1) {
		if currNode == nil {
			return nil, nil
		}

		prevNode, currNode = dp(currNode, currNode.Next)
		counter += 1
		if counter == n {
			prevNode.Next = currNode.Next
		}
		return prevNode, currNode.Next
	}

	prevNode, currNode = dp(prevNode, currNode)
	return currNode
}
