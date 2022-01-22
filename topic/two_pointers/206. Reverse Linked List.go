package two_pointers

import "github.com/InterviewTips/algorithm-coding/guns"

type ListNode = guns.ListNode

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var next *ListNode

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}