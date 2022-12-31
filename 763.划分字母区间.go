/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */
package main

import "fmt"

// @lc code=start
func partitionLabels(S string) []int {
	localtion := make(map[string][]int, len(S))

	for i := 0; i < len(S); i++ {
		localtion[string(S[i])] = append(localtion[string(S[i])], i)
	}

	for i := 0; i < len(map); i++ {}

	fmt.Println(localtion)
	return []int{3}
}

// @lc code=end
