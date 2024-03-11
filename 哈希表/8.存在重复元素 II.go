package main

import "fmt"

/*
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。

示例 1：

输入：nums = [1,2,3,1], k = 3
输出：true
示例 2：

输入：nums = [1,0,1,1], k = 1
输出：true
示例 3：

输入：nums = [1,2,3,1,2,3], k = 2
输出：false

提示：

1 <= nums.length <= 105
-109 <= nums[i] <= 109
0 <= k <= 105

https://leetcode.cn/problems/contains-duplicate-ii/description/
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	num2index := make(map[int]int, len(nums))
	for i, x := range nums {
		if v, ok := num2index[x]; ok {
			if i-v <= k {
				return true
			}
		}
		num2index[x] = i
	}
	return false
}

func main() {
	nums, k := []int{1, 2, 3, 1}, 3
	res := containsNearbyDuplicate(nums, k)
	fmt.Println(res)
}
