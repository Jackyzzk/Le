package main

/*
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

示例 1：

输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]
示例 2：

输入：head = [0,1,2], k = 4
输出：[2,0,1]

提示：

链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 109

https://leetcode.cn/problems/rotate-list/description/
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	var n int
	for h := head; h != nil; h = h.Next {
		n++
		if h.Next == nil {
			h.Next = head
			break
		}
	}

	k %= n
	h := head

	for i := 0; i < n-k-1; i++ {
		h = h.Next
	}
	ret := h.Next
	h.Next = nil
	return ret

}
