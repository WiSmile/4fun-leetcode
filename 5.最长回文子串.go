/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	var (
		reskey int
		reslen int
		sep byte = byte('*')
	)

	slen := len(s)
	fillrune := make([]byte, slen * 2)
	fillrune[0] = s[0]
	for i := 1; i < slen; i++ {
		fillrune[2 * i] = s[i]
		fillrune[2 * i - 1] = sep
	}

	for i, fslen := 0, len(fillrune); i < fslen; i++ {
		// 以该字符为中心拓展
		j := 1
		for {
			if i < j {
				break
			}

			if i + j >= fslen {
				break
			}

			if fillrune[i - j] != fillrune[i + j] {
				break
			}

			j++
		}

		j--
		if j > reslen || (j == reslen && fillrune[i - j] != sep) {
			reslen = j
			reskey = i
		}
	}

	fillres := fillrune[reskey - reslen : reskey + reslen + 1]
	resrune := []byte{}
	for i, j, frlen := 0, 0, len(fillres); i < frlen; i++ {
		if fillres[i] == sep {
			continue
		}

		resrune = append(resrune, fillres[i])
		j++
	}

	return string(resrune)
}
// @lc code=end

