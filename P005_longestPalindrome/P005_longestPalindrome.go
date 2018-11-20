package P005_longestPalindrome

/**
给定一个字符串 s，找到 s 中最长的回文子串。
你可以假设 s 的最大长度为1000。
 */

/**
思路：
	首先想到的是找到所有子串中最长的，但是时间复杂度太高了
	那么就是以每个点为中心扩散最后找到最长的回文串
 */
func longestPalindrome(s string) string {
	if s == `` || len(s) == 0 {
		return ``
	}
	strs := stringToCharArray(s)

	maxlength := 1
	start := 0
	end := 0
	for i, _ := range strs {
		length := 1
		j, k := i-1, i+1
		for j >= 0 && k <= len(strs)-1 && strs[j] == strs[k] {
			j--
			k++
			length ++
		}

		if length > maxlength {
			maxlength = length
			start = j + 1
			end = k - 1
		}

	}

	finishingProduct := strs[start : end+1]
	res := ""
	for _, element := range finishingProduct {
		if element != "#" {
			res += element
		}
	}
	return res
}

func stringToCharArray(str string) (strs []string) {
	for _, element := range str {
		strs = append(append(strs, string(element)), "#")
	}
	return
}
