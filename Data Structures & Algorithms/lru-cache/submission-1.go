type LRUCache struct {
	cache map[int]*ListNode
	size int
	left *ListNode
	right *ListNode
}

type ListNode struct {
    Key int
	Val int
	Next *ListNode
	Prev *ListNode
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
		cache : map[int]*ListNode{},
		size : capacity,
		left : &ListNode{},
		right : &ListNode{},
	}
	lru.right.Prev = lru.left
	lru.left.Next = lru.right
	return lru
}

func (this *LRUCache) remove(node *ListNode) {
	next, prev := node.Next, node.Prev
	next.Prev = prev
	prev.Next = next
}

func (this *LRUCache) insert(node *ListNode) {
	next, prev := this.right, this.right.Prev
	prev.Next = node
	node.Prev = prev
	node.Next = next
	next.Prev = node
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
