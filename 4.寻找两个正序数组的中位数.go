/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		temp := nums1
		nums1 = nums2
		nums2 = temp
	}

	len1 := len(nums1)
	len2 := len(nums2)
	leftCnt := (len1 + len2 + 1) / 2

	left := 0
	right := len1
	i, j := 0, 0
	for left < right {
		i = (left + right) / 2
		j = leftCnt - i - 1
		if nums1[i] > nums2[j] {
			right = i
		} else {
			left = i + 1
		}
	}

	var left1, left2, right1, right2 int
	if left == 0 {
		left1 = math.MinInt32
	} else {
		left1 = nums1[left - 1]
	}

	if leftCnt - left == 0 {
		left2 = math.MinInt32
	} else {
		left2 = nums2[leftCnt - left - 1]
	}

	if len1 == left {
		right1 = math.MaxInt32
	} else {
		right1 = nums1[left]
	}

	if leftCnt - left == len2 {
		right2 = math.MaxInt32
	} else {
		right2 = nums2[leftCnt - left]
	}

	if (len1 + len2) % 2 == 0 {
		return (float64(max(left1, left2)) + float64(min(right1, right2))) / 2
	} else {
		return float64(max(left1, left2))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

