package P004_findMedianSortedArrays

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_ok(t *testing.T) {
	ast := assert.New(t)
	questions := []question{
		question{
			p: param{
				one: []int{1, 3},
				two: []int{2},
			},
			r: result{
				ans: 2,
			},
		},

		question{
			p: param{
				one: []int{1, 2},
				two: []int{3, 4},
			},
			r: result{
				ans: 2.5,
			},
		},
	}
	for _, que := range questions {
		answer, para := que.r, que.p
		ast.Equal(answer.ans, findMedianSortedArrays(para.one, para.two), "case one")
		ast.Equal(answer.ans, findMedianSortedArrays_II(para.one, para.two), "case two")
	}
}

type param struct {
	one []int
	two []int
}

type result struct {
	ans float64
}

type question struct {
	p param
	r result
}
