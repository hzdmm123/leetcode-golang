package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//indexValMap := map[int]int{}
	var indexValMap = make(map[int]int, len(nums))
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
