/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	jinwei := 0
	sum := 0
	temp := result

	for {
		if l1 == nil && l2 == nil {
			break
		} else if l1 == nil {
			sum = l2.Val + jinwei
		} else if l2 == nil {
			sum = l1.Val + jinwei
		} else {
			sum = l1.Val + l2.Val + jinwei
		}

		next := new(ListNode)
		next.Val = sum % 10
		temp.Next = next
		temp = next
		jinwei = sum / 10 % 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if jinwei != 0 {
		temp.Next = &ListNode{Val: jinwei}
	}

	return result.Next
}

// @lc code=end

