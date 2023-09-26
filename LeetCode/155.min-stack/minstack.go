package minstack

import "fmt"

type Node struct {
	val  int
	next *Node
}

type MinStack struct {
	top *Node
}

func Constructor() MinStack {
	return MinStack{nil}
}

func (this *MinStack) Push(val int) {
	this.top = &Node{val, this.top}
}

func (this *MinStack) Pop() {
	this.top = this.top.next
}

func (this *MinStack) Top() int {
	return this.top.val
}

func (this *MinStack) GetMin() int {
	cur := this.top.next
	min := this.top.val
	for cur != nil {
		if cur.val < min {
			min = cur.val
		}
		cur = cur.next
	}
	return min
}

func Main() {
	obj := Constructor()
	obj.Push(3)
	obj.Push(5)
	obj.Push(2)
	obj.Push(5)
	obj.Push(10)
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
}
