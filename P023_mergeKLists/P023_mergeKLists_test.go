package P023_mergeKLists

import (
	"leetcode-golang/kit"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_ok(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		/*		question{
					param{
						[][]int{
							[]int{1, 4, 7},
							[]int{2, 5, 8},
							[]int{3, 6, 9},
						},
					},
				},*/

		question{
			param{
				[][]int{
					{1, 4, 7},
					{2, 5, 8},
					{3, 6, 9},
				}},
			ans{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
	}

	for _, q := range qs {
		answer, para := q.ans, q.param
		ast.Equal(answer.result, List2Slice(mergeKLists(Slice2D2ListNode(para.param))), "terrible")
	}
}

var node1 *kit.ListNode

type param struct {
	param [][]int
}

type ans struct {
	result []int
}

type question struct {
	param
	ans ans
}

func List2Slice(head *kit.ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func Slice2List(nums []int) *kit.ListNode {
	/*
	if len(s) == 0 {
		return nil
	}

	head := &kit.ListNode{
		Val: s[0],
	}
	fake := &kit.ListNode{}
	fake.Next = head

	for i := 1; i < len(s); i++ {
		head.Next = &kit.ListNode{
			Val: s[i],
		}
		head = head.Next
	}
	return fake.Next
	*/

	if len(nums) == 0 {
		return nil
	}

	res := &kit.ListNode{
		Val: nums[0],
	}
	fake := &kit.ListNode{}
	fake.Next = res
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &kit.ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return fake.Next
}

func Slice2D2ListNode(numss [][]int) []*kit.ListNode {
	res := make([]*kit.ListNode, 0)

	for _, num := range numss {
		res = append(res, Slice2List(num))
	}
	return res
}
