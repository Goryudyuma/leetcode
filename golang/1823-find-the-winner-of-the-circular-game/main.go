package main

/*
 * @lc app=leetcode id=1823 lang=golang
 *
 * [1823] Find the Winner of the Circular Game
 */

// @lc code=start
func solve2(array []int, k int) int {
	if len(array) == 1 {
		return array[0]
	}
	i := (k - 1) % len(array)
	return solve2(append(array[i+1:], array[:i]...), k)
}

func findTheWinner(n int, k int) int {
	array := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		array = append(array, i)
	}
	return solve2(array, k)
}

// @lc code=end
