package main

/*
给你一个二叉树的根节点 root ， 检查它是否轴对称。

示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：

输入：root = [1,2,2,null,3,null,3]
输出：false

提示：

树中节点数目在范围 [1, 1000] 内
-100 <= Node.val <= 100

进阶：你可以运用递归和迭代两种方法解决这个问题吗？

https://leetcode.cn/problems/invert-binary-tree/description/
*/

func isSymmetricV1(root *TreeNode) bool {
	var check func(p, q *TreeNode) bool
	check = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return check(p.Left, q.Right) && check(p.Right, q.Left)
	}
	return check(root, root)
}

func isSymmetricV2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) >= 2 {
		p, q := queue[0], queue[1]
		queue = queue[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(queue, p.Left, q.Right, p.Right, q.Left)
	}
	return true
}
