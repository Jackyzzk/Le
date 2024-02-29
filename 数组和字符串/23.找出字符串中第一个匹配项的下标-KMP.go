package main

import "fmt"

/*
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。

示例 1：

输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。
示例 2：

输入：haystack = "leetcode", needle = "leeto"
输出：-1
解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。

提示：

1 <= haystack.length, needle.length <= 104
haystack 和 needle 仅由小写英文字符组成

https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
*/

func strStr(haystack string, needle string) int {
	index, i, j := -1, 0, 0
	n1, n2 := len(haystack), len(needle)
	for i < n1 && j < n2 {
		if haystack[i] != needle[j] {
			i++
			j = 0
			continue
		}
		if j == 0 {
			index = i
		}
		i++
		j++
	}
	if j == n2 {
		return index
	}
	return -1
}

func main() {
	//haystack, needle := "sadbutsad", "sad"
	//haystack, needle := "leetcode", "leeto"
	haystack, needle := "mississippi", "issip" // 4
	res := strStr(haystack, needle)
	fmt.Println(res)
}
