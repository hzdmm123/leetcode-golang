package P015_threeSum

import (
	"sort"
)

/**
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
 */

 //先将所有的数字排序，然后
func threeSum(nums []int) [][]int {
	res := [][]int{}
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			for k:=j+1; k<len(nums); k++ {
				if nums[i]+nums[j]+nums[k]==0 {
					res = append(res,[]int{nums[i],nums[j],nums[k]})
				}
			}
		}
	}
	return res
}

func threeSum_II(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		if i>0 && nums[i] == nums[i-1] {
			continue
		}

		l,r := i+1,len(nums)-1

		for l<r {
			a := nums[i] + nums[l] + nums[r]
			switch {
			case a <0:
				l++
			case a>0:
				r--
			default:
				res = append(res,[]int{nums[i],nums[l],nums[r]})
				l,r = next(nums,l,r)
			}
		}
	}
	return res
}

func next(nums []int,l int,r int)(int,int) {
	for l<r  {
		switch  {
		case nums[l]==nums[l+1]:
			l++
		case nums[r] == nums[r-1]:
			r--
		default:
			l++
			r--
			return l,r
		}
	}
	return l,r
}