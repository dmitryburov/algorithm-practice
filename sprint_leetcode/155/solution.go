package main

type MinStack struct {
	stack, min []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.min) > 0 {
		min := this.GetMin()
		if min < val {
			val = min
		}
	}

	this.min = append(this.min, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

func (this *MinStack) Empty() bool {
	return len(this.stack) > 0
}
