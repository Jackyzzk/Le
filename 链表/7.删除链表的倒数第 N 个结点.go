package main

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]

提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

进阶：你能尝试使用一趟扫描实现吗？

https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ret := &ListNode{Next: head}

	slow, fast := ret, ret
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}

	// fast 领先 slow n个节点, 当 fast 到末尾, slow 就是倒数的那个节点

	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	slow.Next = slow.Next.Next

	return ret.Next
}
