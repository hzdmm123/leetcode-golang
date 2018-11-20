package P005_longestPalindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPalindrome(t *testing.T){
	ast := assert.New(t)

	q := []question{
		question{
			param{
				s:"1abccba2",
			},
			ans{
				s:"abccba",
			},
		},
	}

	for _,element := range q {
		ast.Equal(element.a.s,longestPalindrome(element.p.s))
	}
}

type param struct {
	s string
}

type ans struct {
	s string
}

type question struct {
	p param
	a ans
}

