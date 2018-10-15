package P007

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	result := 0
	var tail int
	var newResult int
	for x != 0 {
		tail = x % 10
		newResult = result*10 + tail
		x = x / 10
		result = newResult
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

func reverseII(x int) int {
	str := strconv.Itoa(x);
	sign := ""
	res := ""
	for _, v := range str {
		char := string(v)
		if char == "-" {
			sign = char
			continue
		}

		res = char + res
	}
	res = sign + res
	resNum, _ := strconv.Atoi(res)
	if resNum > math.MaxInt32 || resNum < math.MinInt32 {
		return 0
	}
	return resNum
}
