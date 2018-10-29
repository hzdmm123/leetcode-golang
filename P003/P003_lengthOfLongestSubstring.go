package P003

import "math"

func lengthOfLongestSubstring(s string) int {
	//类似的一个滑动窗口的概念
	data := []byte(s)
	ans, start, end := 0, 0, 0
	subs := make(map[byte]int)
	for start < len(data) && end < len(data) {
		if _, ok := subs[data[end]]; !ok {
			subs[data[end]] = 1
			end++
			ans = int(math.Max(float64(end-start), float64(ans)))
		} else {
			delete(subs, data[start])
			start++
		}
	}
	return ans
}
