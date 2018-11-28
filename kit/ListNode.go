package kit

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateListNode(length int) *ListNode {
	head := &ListNode{}
	fake := head
	for length > 0 {
		tmp := &ListNode{
			Val:int(rand.Intn(10)),
		}
		head.Next = tmp
		head = tmp
		length--
	}
	return fake.Next
}

func ListNodeEqual(node1 *ListNode, node2 *ListNode)  bool{
	for node1!= nil && node2 != nil {
		if node1.Val!= node2.Val {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}

	if node1 != nil || node2!= nil {
		return false
	}
	return true
}