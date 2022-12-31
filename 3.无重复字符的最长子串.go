/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	input := []rune(s)
	
	head := 0
	length := 1
	repeatLen := 0
	maxSlice := input[head:length]

	for i := 1; i < len(input); i++ {
		repeatLen--
		temp := repeatKey(input[i], maxSlice)
		if temp > repeatLen {
			repeatLen = temp
		}
		
		if repeatLen > 0 {
			head++
		}
		
		length++
		maxSlice = input[head:length]
	}

	return len(maxSlice)
}

func repeatKey(c rune, runeSl []rune) (int) {
	res := 0
	for k, v := range runeSl {
		if c == v {
			res = k + 1
		}
	}

	return res
}

// @lc code=end

