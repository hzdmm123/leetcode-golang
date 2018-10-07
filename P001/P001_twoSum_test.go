package P001

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_ok(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		question{
			param: param{
				one: []int{3, 2, 4},
				two: 6,
			},
			ans: ans{
				res: []int{1, 2},
			},
		},
		question{
			param: param{
				one: []int{3, 2, 4},
				two: 8,
			},
			ans: ans{
				res: nil,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.param
		ast.Equal(a.res, twoSum(p.one, p.two), "输入:%v", p)
	}
}

type param struct {
	one []int
	two int
}

type ans struct {
	res []int
}

type question struct {
	param param
	ans   ans
}
