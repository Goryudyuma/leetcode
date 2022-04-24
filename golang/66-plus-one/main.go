package main

/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	f := true
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			f = false
			digits[i]++
			break
		} else {
			digits[i] = 0
		}
	}
	if f {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// @lc code=end
