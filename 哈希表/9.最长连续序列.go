package main

import "fmt"

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9

提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109

https://leetcode.cn/problems/longest-consecutive-sequence/description/
*/

func longestConsecutive(nums []int) int {
	num2cnt := make(map[int]int, len(nums))
	var ret int
	for _, x := range nums {
		if _, ok := num2cnt[x]; ok {
			continue
		}
		length := 1
		// 如果 x-1存在，在x来之前，x-1一定是右端点
		if v, ok := num2cnt[x-1]; ok {
			length += v
		}
		// 如果 x+1存在，在x来之前，x+1一定是左端点
		if v, ok := num2cnt[x+1]; ok {
			length += v
		}
		// l = right - left + 1
		if v, ok := num2cnt[x-1]; ok {
			// left = right + 1 - l
			num2cnt[x-v] = length
		}
		if v, ok := num2cnt[x+1]; ok {
			// right = l + left - 1
			num2cnt[v+x] = length
		}
		num2cnt[x] = length
		ret = max(ret, length)
	}
	return ret
}

func main() {
	//nums := []int{100, 4, 200, 1, 3, 2}
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	res := longestConsecutive(nums)
	fmt.Println(res)
}
