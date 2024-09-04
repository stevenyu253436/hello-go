type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyStack {
	return MyStack{
		queue1: []int{},
		queue2: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.queue1 = append(this.queue1, x)
}

func (this *MyStack) Pop() int {
	// 将 queue1 中的元素除了最后一个，全部转移到 queue2
	for len(this.queue1) > 1 {
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}

	// 取出 queue1 最后的元素作为栈顶元素
	top := this.queue1[0]
	this.queue1 = this.queue1[1:]

	// 交换 queue1 和 queue2
	this.queue1, this.queue2 = this.queue2, this.queue1

	return top
}

func (this *MyStack) Top() int {
	// 与 Pop 类似，但不移除最后一个元素
	for len(this.queue1) > 1 {
		this.queue2 = append(this.queue2, this.queue1[0])
		this.queue1 = this.queue1[1:]
	}

	top := this.queue1[0]
	this.queue2 = append(this.queue2, this.queue1[0]) // 把栈顶元素也转移到 queue2
	this.queue1 = this.queue1[1:]

	this.queue1, this.queue2 = this.queue2, this.queue1

	return top
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */