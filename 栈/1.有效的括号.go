package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false

提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成

https://leetcode.cn/problems/valid-parentheses/description/
*/

type stack[T any] struct {
	list []T
	n    int
}

func (z *stack[T]) Add(in T) {
	z.list = append(z.list, in)
	z.n++
}

func (z *stack[T]) Pop() (T, bool) {
	if z.n == 0 {
		var t T
		return t, false
	}
	z.n--
	ret := z.list[z.n]
	z.list = z.list[:z.n]
	return ret, true
}

func isValid(s string) bool {
	info := map[byte]byte{
		'(': ')', '{': '}', '[': ']',
	}
	z := stack[byte]{}
	for i := range s {
		x := s[i]
		if _, ok := info[x]; ok {
			z.Add(x)
			continue
		}
		pre, ok := z.Pop()
		if !ok {
			return false
		}
		if x != info[pre] {
			return false
		}
	}
	return z.n == 0
}

func main() {
	s := "()[]{}"
	//s := "([)]"
	//s := "["
	res := isValid(s)
	fmt.Println(res)
}
