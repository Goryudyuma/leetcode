package main

import "sort"

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	type number struct {
		num   int
		index int
	}
	array := make([]number, 0, len(nums))
	for i, num := range nums {
		array = append(array, number{num: num, index: i})
	}
	sort.Slice(array, func(i, j int) bool {
		return array[i].num < array[j].num
	})
	i, j := 0, len(nums)-1
	for i < j {
		if array[i].num+array[j].num < target {
			i++
		} else if array[i].num+array[j].num > target {
			j--
		} else {
			return []int{array[i].index, array[j].index}
		}
	}
	return nil
}

// @lc code=end
