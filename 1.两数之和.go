/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int, len(nums))

	for i, x := range nums {
		value, exist := hashmap[target-x]
		if exist {
			return []int{i, value}
		}
		hashmap[x] = i
	}

	return nil
}

// @lc code=end

