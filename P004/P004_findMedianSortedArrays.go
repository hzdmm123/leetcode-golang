package P004

import (
	"fmt"
	"math"
	"leetcode-golang/kit"
)

/**
题目
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2 。

请找出这两个有序数组的中位数。要求算法的时间复杂度为 O(log (m+n)) 。

你可以假设 nums1 和 nums2 不同时为空。


示例 1:

nums1 = [1, 3]
nums2 = [2]

中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

中位数是 (2 + 3)/2 = 2.5
 */

/**
使用二分在两个排序数组找到中间那个数字，时间复杂度  o(m+n)
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//简单的方式是开个新的数组依次放入num的值
	slice := make([]int, 0)
	i, j := 0, 0
	for ; i < len(nums1) && j < len(nums2);
	{
		if nums1[i] < nums2[j] {
			slice = append(slice, nums1[i])
			i++
		} else {
			slice = append(slice, nums2[j])
			j++
		}
	}

	if i == len(nums1) {
		for ; j < len(nums2); j++ {
			slice = append(slice, nums2[j])
		}
	}

	if j == len(nums2) {
		fmt.Printf("this is j=[%v]", j)
		for ; i < len(nums1); i++ {
			slice = append(slice, nums1[i])
		}
	}

	if len(slice)%2 == 1 {
		return float64(slice[len(slice)/2])
	} else {
		return (float64(slice[len(slice)/2-1]) + float64(slice[len(slice)/2])) / 2
	}
}

func findMedianSortedArrays_II(nums1 []int, nums2 []int) float64 {
	if len(nums2) < len(nums1) {
		findMedianSortedArrays_II(nums2, nums1)
	}
	l1, l2, c1, r1, r2, c2 := 0, 0, 0, 0, 0, 0
	low, high := 0, 2*len(nums1)
	for low <= high {
		c1 = (low + high) / 2
		c2 = len(nums1) + len(nums2) - c1
		l1 = kit.If(c1 == 0, math.MinInt32, nums1[int(c1-1/2)]).(int)
		r1 = kit.If(c1 == 2*len(nums1), math.MaxInt32, nums1[int(c1/2)]).(int)
		l2 = kit.If(c2 == 0, math.MinInt32, nums2[int((c2-1)/2)]).(int)
		r2 = kit.If(c2 == 2*len(nums2), math.MaxInt32, nums2[int((c2)/2)]).(int)
		if l1 > r2 {
			high = c1 - 1
		} else if l2 > r1 {
			low = c1 + 1
		} else {
			break
		}
	}
	return (math.Max(float64(l1), float64(l2)) + math.Min(float64(r1), float64(r2))) / 2.0
}
