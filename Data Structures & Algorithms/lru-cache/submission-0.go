type ListNode struct {
	Key int
	Val int
	Next *ListNode
	Prev *ListNode
}

type LRUCache struct {
    cache map[int]*ListNode
	size int
	left *ListNode
	right *ListNode
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{cache : map[int]*ListNode{}, size: capacity,
	left : &ListNode{}, right: &ListNode{}}
	lru.left.Next = lru.right
	lru.right.Prev = lru.left
	return lru
}

func (this *LRUCache) remove(node *ListNode) {
	prev, next := node.Prev, node.Next
    prev.Next = next
    next.Prev = prev
}

func (this *LRUCache) insert(node *ListNode) {
	prev, next := this.right.Prev, this.right
	prev.Next = node
	next.Prev = node
	node.Next = next
	node.Prev = prev
}

func (this *LRUCache) Get(key int) int {
    if node, exists := this.cache[key]; exists {
		this.remove(node)
		this.insert(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exists := this.cache[key]; exists {
		this.remove(node)
	}
	this.cache[key] = &ListNode{Key: key, Val: value}
	this.insert(this.cache[key])

	if len(this.cache) > this.size {
		lru := this.left.Next
		this.remove(lru)
		delete(this.cache, lru.Key)
	}
}
