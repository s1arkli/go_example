package main

/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/

type node struct {
	key  int
	val  int
	pre  *node
	next *node
}

type LRUCache struct {
	cap  int
	m    map[int]*node
	head *node
	tail *node
}

func Constructor(capacity int) LRUCache {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		cap:  capacity,
		m:    make(map[int]*node, capacity),
		head: head,
		tail: tail,
	}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.m[key]; ok {
		c.remove(node)
		c.pushFront(node)
		return node.val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.m[key]; ok {
		node.val = value
		c.remove(node)
		c.pushFront(node)
		return
	}
	if c.cap == len(c.m) {
		delete(c.m, c.tail.pre.key)
		c.remove(c.tail.pre)
	}
	newNode := &node{key, value, nil, nil}
	c.m[key] = newNode
	c.pushFront(newNode)
}

func (c *LRUCache) remove(l *node) {
	l.pre.next = l.next
	l.next.pre = l.pre
}

func (c *LRUCache) pushFront(l *node) {
	right := c.head.next
	l.next = right
	l.pre = c.head
	c.head.next = l
	right.pre = l
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
