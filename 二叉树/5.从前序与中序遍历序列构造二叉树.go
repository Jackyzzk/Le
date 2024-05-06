package main

/*
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

示例 1:

输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]
示例 2:

输入: preorder = [-1], inorder = [-1]
输出: [-1]

提示:

1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder 和 inorder 均 无重复 元素
inorder 均出现在 preorder
preorder 保证 为二叉树的前序遍历序列
inorder 保证 为二叉树的中序遍历序列

先序遍历也叫做前序遍历、先根遍历，可记做根左右。
中序遍历 左根右
后序遍历 左右根

这里的序说的是根在哪

https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
*/

// [ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
// [ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]

func buildTreeV1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; ; i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	root.Left = buildTreeV1(preorder[1:i+1], inorder[:i])
	root.Right = buildTreeV1(preorder[i+1:], inorder[i+1:])
	return root
}

// 优化每次找根节点的时间消耗
func buildTreeV2(preorder []int, inorder []int) *TreeNode {
	n := len(inorder)
	root2index := make(map[int]int, n)
	for i, x := range inorder {
		root2index[x] = i
	}

	var build func(p1, p2, i1, i2 int) *TreeNode
	build = func(p1, p2, i1, i2 int) *TreeNode {
		if p1 > p2 || i1 > i2 {
			return nil
		}
		root := &TreeNode{Val: preorder[p1]}
		index := root2index[root.Val]
		// 注意这个index是inorder的 不可直接用于preorder
		// 在preorder使用时需要转换为步长
		n1, n2 := index-i1-1, i2-index-1
		root.Left = build(p1+1, p1+1+n1, i1, index-1)
		root.Right = build(p2-n2, p2, index+1, i2)
		return root
	}

	return build(0, n-1, 0, n-1)
}
