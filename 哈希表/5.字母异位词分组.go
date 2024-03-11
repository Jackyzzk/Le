package main

import "fmt"

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]

提示：

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母

https://leetcode.cn/problems/group-anagrams/description/
*/

func groupAnagrams(strs []string) [][]string {
	arr2str := make(map[[26]int][]string, len(strs))
	for _, x := range strs {
		arr := [26]int{}
		for _, r := range x {
			arr[r-'a']++
		}
		arr2str[arr] = append(arr2str[arr], x)
	}
	ret := make([][]string, 0, len(arr2str))
	for _, x := range arr2str {
		ret = append(ret, x)
	}
	return ret
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}
