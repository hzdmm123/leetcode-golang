package P009_isPalindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		question{
			param{
				s: 121,
			},
			ans{
				result: true,
			},
		},
		question{
			param{
				s: -121,
			},
			ans{
				result: false,
			},
		},
	}

	for _, qss := range qs {
		ast.Equal(qss.a.result, isPalindrome(qss.p.s))
	}
}

type param struct {
	s int
}

type ans struct {
	result bool
}

type question struct {
	p param
	a ans
}
