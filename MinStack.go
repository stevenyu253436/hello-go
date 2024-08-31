package main

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt32}, // 初始化 minStack，以確保 getMin 不會出錯
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	// 將當前值與 minStack 的最小值比較，將較小者壓入 minStack
	minVal := this.minStack[len(this.minStack)-1]
	if val < minVal {
		minVal = val
	}
	this.minStack = append(this.minStack, minVal)
}

func (this *MinStack) Pop() {
	// 移除 stack 和 minStack 的頂部元素
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func main() {
	// 這裡可以寫測試代碼
}
