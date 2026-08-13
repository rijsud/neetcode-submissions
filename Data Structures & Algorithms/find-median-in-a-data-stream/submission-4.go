type MinHeap []int
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int))}
func (h *MinHeap) Pop() interface{} {
old := *h
n := len(old)
pop := old[n-1]
*h = old[:n-1]
return pop
}

type MaxHeap []int
func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int))}
func (h *MaxHeap) Pop() interface{} {
old := *h
n := len(old)
pop := old[n-1]
*h = old[:n-1]
return pop
}

// Less for MinHeap (Smallest element climbs to the top)
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
// Less for MaxHeap (Largest element climbs to the top)
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }

type MedianFinder struct {
    large *MinHeap
	small *MaxHeap
}


func Constructor() MedianFinder {
    medFind := MedianFinder {
		large : &MinHeap{},
		small : &MaxHeap{},
	}
	heap.Init(medFind.large)
	heap.Init(medFind.small)
	return medFind
}


func (this *MedianFinder) AddNum(num int)  {
	if this.large.Len() == 0 || num > (*this.large)[0] {
		heap.Push(this.large, num)
	} else {
		heap.Push(this.small, num)
	}

    if this.large.Len() > this.small.Len() + 1 {
		heap.Push(this.small, heap.Pop(this.large))
	} else if this.small.Len() > this.large.Len() + 1 {
		heap.Push(this.large, heap.Pop(this.small))
	}
}


// func (this *MedianFinder) FindMedian() float64 {
//     n := this.small.Len() + this.large.Len()
// 	var res float64
// 	var1 := (*this.large)[0]
// 	var2 := (*this.small)[0]
// 	if n % 2 != 0 {
// 		if this.large.Len() > this.small.Len() {
// 			res = float64(var1)
// 		} else {
// 			res = float64(var2)
// 		}
// 	} else {
// 		res = float64(var1 + var2) / 2.0
// 	}
// 	return res
// }

func (this *MedianFinder) FindMedian() float64 {
    if this.small.Len() > this.large.Len() {
        return float64((*this.small)[0])
    }
    if this.large.Len() > this.small.Len() {
        return float64((*this.large)[0])
    }
    return float64((*this.small)[0]+(*this.large)[0]) / 2.0
}
