package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

提示：

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成

https://leetcode.cn/problems/longest-common-prefix/description/
*/

func longestCommonPrefix(strs []string) string {
	ret, n := -1, len(strs)
	for i := range strs[0] {
		x := strs[0][i]
		ok := true
		for j := 1; j < n; j++ {
			if i >= len(strs[j]) || strs[j][i] != x {
				ok = false
				break
			}
		}
		if ok {
			ret = i
		} else {
			break
		}
	}
	return strs[0][:ret+1]
}

func main() {
	//strs := []string{"flower", "flow", "flight"}
	//strs := []string{"dog", "racecar", "car"}
	strs := []string{"cir", "car"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}
