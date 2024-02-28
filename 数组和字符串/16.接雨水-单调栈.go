package main

import "fmt"

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9

提示：

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105

https://leetcode.cn/problems/trapping-rain-water/description/
*/

type Z struct {
	list []int
	n    int
}

func (z *Z) Add(in int) {
	z.list = append(z.list, in)
	z.n++
}

func (z *Z) Last() (int, bool) {
	if z.n <= 0 {
		return 0, false
	}
	return z.list[z.n-1], true
}

func (z *Z) Pop() (int, bool) {
	if z.n <= 0 {
		return 0, false
	}
	z.n--
	ret := z.list[z.n]
	z.list = z.list[:z.n]
	return ret, true
}

func trap(height []int) int {
	z := Z{}
	ret := 0
	for i, x := range height {
		for {
			last, ok := z.Last()
			if !ok || height[last] >= x {
				z.Add(i)
				break
			}
			mid, _ := z.Pop()
			left, ok := z.Last()
			if ok {
				ret += (min(height[left], x) - height[mid]) * (i - left - 1)
			}
		}
	}
	return ret
}

func main() {
	//height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//height := []int{4, 2, 0, 3, 2, 5}
	height := []int{5, 2, 1, 2, 1, 5} // 14
	res := trap(height)
	fmt.Println(res)
}
