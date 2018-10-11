package P003

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_ok(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(0, reverse(1534236469), "over int")

	questions := []question{
		question{
			p: param{
				in: 12345670,
			},
			a: ans{
				out: 7654321,
			},
		},
	}

	for _, q := range questions {
		answer, p := q.a, q.p
		ast.Equal(answer.out, reverse(p.in), "输入:%v", p)
		ast.Equal(answer.out, reverseII(p.in), "输入:%v", p)
	}
}

type param struct {
	in int
}

type ans struct {
	out int
}

type question struct {
	p param
	a ans
}
