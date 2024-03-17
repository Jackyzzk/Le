package main

import "fmt"

/*
给你一个 无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

示例 1：

输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]
示例 2：

输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
示例 3：

输入：intervals = [], newInterval = [5,7]
输出：[[5,7]]
示例 4：

输入：intervals = [[1,5]], newInterval = [2,3]
输出：[[1,5]]
示例 5：

输入：intervals = [[1,5]], newInterval = [2,7]
输出：[[1,7]]

提示：

0 <= intervals.length <= 104
intervals[i].length == 2
0 <= intervals[i][0] <= intervals[i][1] <= 105
intervals 根据 intervals[i][0] 按 升序 排列
newInterval.length == 2
0 <= newInterval[0] <= newInterval[1] <= 105

https://leetcode.cn/problems/insert-interval/description/
*/

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	index := 0
	for i, x := range intervals {
		if x[0] > newInterval[0] {
			index = i
			intervals = append(intervals[:i], append([][]int{newInterval}, intervals[i:]...)...)
			break
		}
	}
	if index > 0 {
		index--
	}
	ret := intervals[:index]

	l, r := intervals[index][0], intervals[index][1]
	for i := index + 1; i < len(intervals); i++ {
		x, y := intervals[i][0], intervals[i][1]
		if x <= r && y > r {
			r = y
			continue
		}
		if x > r {
			index = i
			break
		}
	}
	ret = append(ret, append([][]int{{l, r}}, intervals[index:]...)...)
	return ret
}

func main() {
	//intervals, newInterval := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}
	intervals, newInterval := [][]int{{1, 5}}, []int{2, 3} // [[1,5]]
	res := insert(intervals, newInterval)
	fmt.Println(res)
}
