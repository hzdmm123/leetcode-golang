package P023_mergeKLists

import "leetcode-golang/kit"

/**
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
 */
func mergeKLists(lists []*kit.ListNode) *kit.ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists)
}

func merge(list []*kit.ListNode) *kit.ListNode {
	if len(list) == 1 {
		return merge2Lists(list[0], nil)
	}

	if len(list) == 2 {
		return merge2Lists(list[0], list[1])
	}

	half := len(list) / 2
	return merge([]*kit.ListNode{merge(list[:half]), merge(list[half:])})
}

//先解决两个listmerge的问题
func merge2Lists(list1 *kit.ListNode, list2 *kit.ListNode) *kit.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	//先找个head
	var res, cur *kit.ListNode
	if list1.Val < list2.Val {
		res, cur, list1 = list1, list1, list1.Next
	} else {
		res, cur, list2 = list2, list2, list2.Next
	}

	//head后面的
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next, list1 = list1, list1.Next
		} else {
			cur.Next, list2 = list2, list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	}

	if list2 != nil {
		cur.Next = list2
	}

	return res
}
