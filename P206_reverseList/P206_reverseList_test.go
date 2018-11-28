package P206_reverseList

import (
	"github.com/stretchr/testify/assert"
	"leetcode-golang/kit"
	"testing"
)

func TestReverseList(t *testing.T) {
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
				kit.GenerateListNode(5)},
			ans{true},
		},
	}

	for _, q := range qs {
		answer, para := q.ans, q.param
		ast.Equal(answer.result,kit.ListNodeEqual(ReverseList(para.ListNode),ReverseList(para.ListNode)))
	}
}

type param struct {
	ListNode *kit.ListNode
}

type ans struct {
	result bool
}

type question struct {
	param
	ans ans
}