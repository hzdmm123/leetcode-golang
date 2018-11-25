package P014_longestCommonPrefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestP014_longestCommonPrefix(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			param{
				strs: []string{"flower","flw","fls"},
			},
			ans{
				result: "fl",
			},
		},
	}


	for _,element := range qs {
		ast.Equal(element.a.result,P014_longestCommonPrefix(element.p.strs))
	}
}


type param struct {
	strs []string
}

type ans struct {
	result string
}

type question struct {
	p param
	a ans
}
