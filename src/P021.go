package src

// 合并两个有序链表
// https://leetcode-cn.com/problems/merge-two-sorted-lists/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var ans *ListNode
	if l1.Val <= l2.Val {
		ans = l1
		l1 = l1.Next
	} else {
		ans = l2
		l2 = l2.Next
	}

	current := ans
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	}

	if l2 != nil {
		current.Next = l2
	}

	return ans
}
