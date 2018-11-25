package P014_longestCommonPrefix

import (
	"math"
)

func P014_longestCommonPrefix(strs []string) string{
	if strs == nil || len(strs) == 0 {
		return ""
	}
	index,maxlen := 0, math.MaxInt8
	for i, element := range  strs {
		length := len(element)
		if length < maxlen {
			maxlen = len(element)
			index = i
		}
	}

	for i,r := range strs[index]{
		for j:=0 ;j<len(strs) ;j++  {
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}
	return  strs[index]
}