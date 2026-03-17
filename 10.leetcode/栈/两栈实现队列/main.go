package main

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Transfer() {
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			top := this.in[len(this.in)-1]
			this.in = this.in[:len(this.in)-1]
			this.out = append(this.out, top)
		}
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	this.Transfer()
	top := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return top
}

func (this *MyQueue) Peek() int {
	this.Transfer()
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.out) == 0 && len(this.in) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
