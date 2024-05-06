package main

/*
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。

示例 1:

输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出：[3,9,20,null,null,15,7]
示例 2:

输入：inorder = [-1], postorder = [-1]
输出：[-1]

提示:

1 <= inorder.length <= 3000
postorder.length == inorder.length
-3000 <= inorder[i], postorder[i] <= 3000
inorder 和 postorder 都由 不同 的值组成
postorder 中每一个值都在 inorder 中
inorder 保证是树的中序遍历
postorder 保证是树的后序遍历

后序遍历（Postorder Traversal）是二叉树遍历的一种形式，其特点是先遍历左子树，然后遍历右子树，最后访问根节点

https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
*/

func buildTreeV3(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	v2index := make(map[int]int, len(inorder))
	for i, x := range inorder {
		v2index[x] = i
	}
	var build func(i1, i2, p1, p2 int) *TreeNode
	build = func(i1, i2, p1, p2 int) *TreeNode {
		if i1 > i2 || p1 > p2 {
			return nil
		}
		root := &TreeNode{Val: postorder[p2]}
		index := v2index[root.Val]
		n1, n2 := index-1-i1, i2-index-1
		root.Left = build(i1, index-1, p1, p1+n1)
		root.Right = build(index+1, i2, p2-1-n2, p2-1)
		return root
	}

	return build(0, n-1, 0, n-1)
}
