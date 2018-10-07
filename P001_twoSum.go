package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//indexValMap := map[int]int{}
	var indexValMap = make(map[int]int, len(nums)) //make的时候确定大小,比没有确定大小的map快很多
	for index, val := range nums {
		if v, ok := indexValMap[target-val]; ok {
			return []int{v, index}
		}
		indexValMap[val] = index
	}
	return nil
}

func main() {
	array := twoSum([]int{1, 2, 4}, 6)
	for index, val := range array {
		fmt.Printf("array[%d]=%d \n", index, val)
	}
}
