package P002

import "leetcode-golang/kit"

/**

给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，
它们的每个节点只存储单个数字。将两数相加返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。


示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
 */
/**
 * Definition for singly-linked list.
**/
type ListNode = kit.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	fakeNode := &ListNode{}
	cur := fakeNode
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return fakeNode.Next
}
