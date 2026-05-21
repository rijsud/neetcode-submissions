type MinStack struct {
	minStack []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, minStack: []int{}}
}

func (this *MinStack) Push(val int) {
	minNumber := val
	if len(this.minStack) != 0 {
		if this.GetMin() < minNumber {minNumber = this.GetMin()}
	}
	this.stack = append(this.stack, val)
	this.minStack = append(this.minStack, minNumber)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
