package P003

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_ok(t *testing.T) {
	ast := assert.New(t)
	questions := []question{
		question{
			p: param{
				str: "bbbbbb",
			},
			a: ans{
				res: 1,
			},
		},

		question{
			p: param{
				str: "abcdeabdij87",
			},
			a: ans{
				res: 8,
			},
		},
	}

	for _, qu := range questions {
		in, out := qu.p.str, qu.a.res
		ast.Equal(out, lengthOfLongestSubstring(in), "ok")
	}
}

type param struct {
	str string
}

type ans struct {
	res int
}

type question struct {
	p param
	a ans
}
