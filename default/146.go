package main

// Design

// 160ms
// 15.1MB
type LRUCache struct {
	data []int
	m    map[int]int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		data: make([]int, capacity),
		m:    make(map[int]int),
	}
	for i := range lru.data {
		lru.data[i] = -1
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	res := -1
	index := -1
	for i, v := range this.data {
		if key == v {
			index = i
			res = this.m[key]
		}
	}
	if index != -1 {
		this.data = append(append(this.data[0:index], this.data[(index + 1):]...), key)
	}
	return res
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) == -1 {
		this.data = append(this.data[1:], key)
	}
	this.m[key] = value
}

// todo:
// 144ms
// 16.1MB
type LRUCache2 struct {
	len, cap    int
	first, last *doubleLinkedList
	m           map[int]*doubleLinkedList
}
type doubleLinkedList struct {
	key, val   int
	next, prev *doubleLinkedList
}

func Constructor2(capacity int) LRUCache2 {
	lru := LRUCache2{
		len:   0,
		cap:   capacity,
		first: nil,
		last:  nil,
		m:     make(map[int]*doubleLinkedList),
	}
	return lru
}

func (this *LRUCache2) Get2(key int) int {
	if node, ok := this.m[key]; ok {
		val := node.val
		this.moveToFirst(node)
		return val
	}
	return -1
}

func (this *LRUCache2) Put2(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.val = value
		this.moveToFirst(node)
	} else {
		node := doubleLinkedList{
			key:  key,
			val:  value,
			prev: nil,
			next: nil,
		}
		this.m[key] = &node
		if this.len == this.cap {
			delete(this.m, this.last.key)
			this.removeLast()
		} else {
			this.len++
		}
		this.insertToFirst(&node)
	}
}

func (this *LRUCache2) removeLast() {
	if this.last.prev != nil {
		this.last.prev.next = nil
		this.last = this.last.prev
	} else {
		this.first = nil
		this.last = nil
	}
}

func (this *LRUCache2) insertToFirst(node *doubleLinkedList) {
	if this.first != nil {
		node.next = this.first
		this.first.prev = node
	} else {
		this.last = node
	}
	this.first = node
}

func (this *LRUCache2) moveToFirst(node *doubleLinkedList) {
	switch node {
	case this.first:
		return
	case this.last:
		this.removeLast()
	default:
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	this.insertToFirst(node)
}
