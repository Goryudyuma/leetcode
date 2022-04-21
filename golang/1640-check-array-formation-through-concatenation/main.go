package main

/*
 * @lc app=leetcode id=1640 lang=golang
 *
 * [1640] Check Array Formation Through Concatenation
 */

// @lc code=start
func solve(arr []int, pieces [][]int) bool {
	if len(arr) == 0 {
		return true
	}
	for _, piece := range pieces {
		if len(piece) > len(arr) {
			continue
		}
		f := true
		for i, v := range piece {
			if arr[i] != v {
				f = false
				break
			}
		}
		if f {
			return solve(arr[len(piece):], pieces)
		}
	}
	return false
}

func canFormArray(arr []int, pieces [][]int) bool {
	return solve(arr, pieces)
}

// @lc code=end
