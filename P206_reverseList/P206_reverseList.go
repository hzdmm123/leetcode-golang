package P206_reverseList

import "leetcode-golang/kit"

func ReverseList(head *kit.ListNode) *kit.ListNode {
	if head == nil {
		return nil
	}
	var  pre *kit.ListNode
	var next *kit.ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
