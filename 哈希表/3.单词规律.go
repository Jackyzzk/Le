package main

import (
	"fmt"
	"strings"
)

/*
给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", s = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", s = "dog cat cat fish"
输出: false
示例 3:

输入: pattern = "aaaa", s = "dog cat cat dog"
输出: false

提示:

1 <= pattern.length <= 300
pattern 只包含小写英文字母
1 <= s.length <= 3000
s 只包含小写英文字母和 ' '
s 不包含 任何前导或尾随对空格
s 中每个单词都被 单个空格 分隔

https://leetcode.cn/problems/word-pattern/description/
*/

func wordPattern(pattern string, s string) bool {
	sList := strings.Split(s, " ")
	n1, n2 := len(pattern), len(sList)
	if n1 != n2 {
		return false
	}
	p2s := make(map[byte]string, n1)
	s2p := make(map[string]byte, n2)
	for i := range pattern {
		if v, ok := p2s[pattern[i]]; !ok {
			p2s[pattern[i]] = sList[i]
		} else if v != sList[i] {
			return false
		}
		if v, ok := s2p[sList[i]]; !ok {
			s2p[sList[i]] = pattern[i]
		} else if v != pattern[i] {
			return false
		}
	}
	return true
}

func main() {
	pattern, s := "abba", "dog cat cat fish"
	res := wordPattern(pattern, s)
	fmt.Println(res)
}
