package kit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringToCharArray(t *testing.T) {
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
				"hello",
			},
			ans{[]string{"h", "e", "l", "l", "0"}},
		},
	}

	for _, q := range qs {
		param, ansower := q.p, q.a
		ast.Equal(len(ansower.s), len(StringToCharArray(param.s)))
	}
}

type param struct {
	s string
}

type ans struct {
	s []string
}

type question struct {
	p param
	a ans
}
