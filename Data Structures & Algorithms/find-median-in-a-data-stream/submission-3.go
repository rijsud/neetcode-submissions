type Heap []int
func (h Heap) Len() int { return len(h) }
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int))}
func (h *Heap) Pop() interface{} {
old := *h
n := len(old)
pop := old[n-1]
*h = old[:n-1]
return pop
}

// 1. MinHeap wraps the base Heap 
type MinHeap struct{ Heap } 
func (h MinHeap) Less(i, j int) bool { return h.Heap[i] < h.Heap[j] } 
// 2. MaxHeap wraps the base Heap 
type MaxHeap struct{ Heap } 
func (h MaxHeap) Less(i, j int) bool { return h.Heap[i] > h.Heap[j] }

type MedianFinder struct {
    large *MinHeap
	small *MaxHeap
}


func Constructor() MedianFinder {
    medFind := MedianFinder {
		large : &MinHeap{Heap: Heap{}},
		small : &MaxHeap{Heap: Heap{}},
	}
	heap.Init(medFind.large)
	heap.Init(medFind.small)
	return medFind
}


func (this *MedianFinder) AddNum(num int)  {
	if this.large.Len() > 0 {
		if num > (*this.large).Heap[0] {
			heap.Push(this.large, num)
		} else {
			heap.Push(this.small, num)
		}
	} else {
		heap.Push(this.small, num)
	}

    if this.large.Len() > this.small.Len() + 1 {
		heap.Push(this.small, heap.Pop(this.large))
	}
	
	if this.small.Len() > this.large.Len() + 1 {
		heap.Push(this.large, heap.Pop(this.small))
	}
}

func (this *MedianFinder) FindMedian() float64 {
    if this.small.Len() > this.large.Len() {
        return float64((*this.small).Heap[0])
    }
    if this.large.Len() > this.small.Len() {
        return float64((*this.large).Heap[0])
    }
    return float64((*this.small).Heap[0]+(*this.large).Heap[0]) / 2.0
}
