/*
 * @lc app=leetcode.cn id=1732 lang=golang
 *
 * [1732] 找到最高海拔
 */

// @lc code=start
func largestAltitude(gain []int) int {
	result := 0
    temp := 0
    for i := 0; i < len(gain); i++ {
        temp = temp + gain[i]
        if temp > result {
            result = temp
        }
    }

    return result
}
// @lc code=end

