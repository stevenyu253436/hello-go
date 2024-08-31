package main

import "fmt"

type LRUCache struct {
	capacity int
	cache    map[int]*DLinkedNode
	head     *DLinkedNode
	tail     *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		node := &DLinkedNode{key: key, value: value}
		this.cache[key] = node
		this.addNode(node)
		if len(this.cache) > this.capacity {
			tail := this.popTail()
			delete(this.cache, tail.key)
		}
	}
}

func (this *LRUCache) addNode(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) popTail() *DLinkedNode {
	res := this.tail.prev
	this.removeNode(res)
	return res
}

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)
	lRUCache.Put(2, 2)
	fmt.Println(lRUCache.Get(1)) // 返回 1
	lRUCache.Put(3, 3)           // 这个操作会使得关键字 2 作废
	fmt.Println(lRUCache.Get(2)) // 返回 -1 (未找到)
	lRUCache.Put(4, 4)           // 这个操作会使得关键字 1 作废
	fmt.Println(lRUCache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(lRUCache.Get(3)) // 返回 3
	fmt.Println(lRUCache.Get(4)) // 返回 4
}
