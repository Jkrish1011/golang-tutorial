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
	lengthOfLinkedList := 0
	/*
		Iterate through the linkedlist to find the length.
		do length - n(nth node from last.)
		Iterate again to reach the (len - n)th node.
		from that node,
			do node.next = node.next.next (this will be as if the next node is not part of the linked list!)

	*/
	for currNode != nil {
		currNode = currNode.Next
		lengthOfLinkedList += 1
	}

	if lengthOfLinkedList == n {
		return head.Next
	}

	currNode = head
	for idx := 0; idx < (lengthOfLinkedList - n - 1); idx++ {
		currNode = currNode.Next
	}

	currNode.Next = currNode.Next.Next

	return head
}
