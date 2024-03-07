package main

import "fmt"

/*
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。

示例 1：

输入：ransomNote = "a", magazine = "b"
输出：false
示例 2：

输入：ransomNote = "aa", magazine = "ab"
输出：false
示例 3：

输入：ransomNote = "aa", magazine = "aab"
输出：true

提示：

1 <= ransomNote.length, magazine.length <= 105
ransomNote 和 magazine 由小写英文字母组成

https://leetcode.cn/problems/ransom-note/description/
*/

func canConstruct(ransomNote string, magazine string) bool {
	byte2int := make(map[byte]int, len(magazine))
	for i := range magazine {
		byte2int[magazine[i]]++
	}
	for i := range ransomNote {
		byte2int[ransomNote[i]]--
		if byte2int[ransomNote[i]] < 0 {
			return false
		}
	}
	return true
}

func main() {
	ransomNote, magazine := "a", "b"
	res := canConstruct(ransomNote, magazine)
	fmt.Println(res)
}
