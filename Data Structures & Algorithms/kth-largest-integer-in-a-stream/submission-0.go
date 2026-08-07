type KthLargest struct {
    Vals []int
	kthVal int
}


func Constructor(k int, nums []int) KthLargest {
    kthVals := KthLargest{
		Vals : nums,
		kthVal : k,
	}
	sort.Slice(kthVals.Vals, func(i, j int) bool {
		return kthVals.Vals[i] > kthVals.Vals[j]
	})
	return kthVals
}


func (this *KthLargest) Add(val int) int {
	this.Vals = append(this.Vals, val)
	sort.Slice(this.Vals, func(i, j int) bool {
		return this.Vals[i] > this.Vals[j]
	})
	return this.Vals[this.kthVal-1]
}
