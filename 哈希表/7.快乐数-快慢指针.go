package main

import "fmt"

/*
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

示例 1：

输入：n = 19
输出：true
解释：
2 是平方运算
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
示例 2：

输入：n = 2
输出：false

提示：

1 <= n <= 231 - 1

https://leetcode.cn/problems/happy-number/description/
*/

func isHappy(n int) bool {
	num2bool := make(map[int]bool)
	for {
		if num2bool[n] {
			return false
		}
		num2bool[n] = true
		var next int
		for n > 0 {
			next += (n % 10) * (n % 10)
			n /= 10
		}
		if next == 1 {
			return true
		}
		n = next
	}
}

// 循环周期问题 快慢指针
func isHappyV2(n int) bool {
	slow, fast := n, n
	fn := func(in int) int {
		var next int
		for in > 0 {
			mod := in % 10
			next += mod * mod
			in /= 10
		}
		return next
	}
	for {
		slow = fn(slow)
		fast = fn(fn(fast))
		if fast == 1 {
			return true
		}
		if slow == fast {
			return false
		}
	}
}

func main() {
	n := 2
	res := isHappyV2(n)
	fmt.Println(res)
}
