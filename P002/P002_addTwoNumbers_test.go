package P002

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_ok(t *testing.T) {
	ast := assert.New(t)
	questions := []question{
		question{
			param: param{
				one: makeListNode([]int{2, 4, 3}),
				two: makeListNode([]int{5, 6, 4}),
			},
			res: ans{
				res: makeListNode([]int{7, 0, 8}),
			},
		},
	}

	for _, q := range questions {
		answer, p := q.res, q.param
		ast.Equal(answer.res, addTwoNumbers(p.one, p.two), "输入:%v", p)
	}

}

func makeListNode(elements []int) *ListNode {
	if elements == nil || len(elements) == 0 {
		return nil
	}

	head := &ListNode{Val: elements[0]}
	cur := head
	for i := 1; i < len(elements); i++ {
		cur.Next = &ListNode{Val: elements[i]}
		cur = cur.Next
	}
	return head
}

type param struct {
	one *ListNode
	two *ListNode
}

type ans struct {
	res *ListNode
}

type question struct {
	param param
	res   ans
}
