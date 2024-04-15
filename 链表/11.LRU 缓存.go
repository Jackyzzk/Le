package main

/*
Least Recently Used
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；
如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

提示：

1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 105
最多调用 2 * 105 次 get 和 put

https://leetcode.cn/problems/lru-cache/description/
*/

type DL struct {
	key   int
	value int
	left  *DL
	right *DL
}

type LRUCache struct {
	n    int
	k2d  map[int]*DL
	top  *DL
	tail *DL
}

func Constructor(capacity int) LRUCache {
	top := &DL{}
	tail := &DL{}
	top.right, tail.left = tail, top
	return LRUCache{
		n:    capacity,
		k2d:  make(map[int]*DL, capacity),
		top:  top,
		tail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	d, ok := this.k2d[key]
	if !ok {
		return -1
	}

	if d != this.top.right {
		d.left.right, d.right.left = d.right, d.left
		d.left, d.right = this.top, this.top.right
		this.top.right.left, this.top.right = d, d
	}

	return d.value
}

func (this *LRUCache) Put(key int, value int) {
	v := this.Get(key)
	if v != -1 {
		this.k2d[key].value = value
		return
	}

	d := &DL{key: key, value: value, left: this.top, right: this.top.right}
	this.top.right.left, this.top.right = d, d

	if len(this.k2d) == this.n {
		delete(this.k2d, this.tail.left.key)
		this.tail.left.left.right, this.tail.left = this.tail, this.tail.left.left
	}
	this.k2d[key] = d
}
