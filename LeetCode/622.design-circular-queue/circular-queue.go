package circularqueue

import "fmt"

type MyCircularQueue struct {
	maxSize int `default:"0"`
	front   int `default:"0"`
	rear    int `default:"-1"`
	q       []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		maxSize: k,
		front:   0,
		rear:    -1,
		q:       make([]int, k, k),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.rear = (this.rear + 1) % this.maxSize
	this.q[this.rear] = value
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.front == this.rear {
		this.front = 0
		this.rear = -1
	} else {
		this.front = (this.front + 1) % this.maxSize
	}
	return true

}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.q[this.front]
	}
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() == true {
		return -1
	} else {
		return this.q[this.rear]
	}

}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.rear == -1
}

func (this *MyCircularQueue) IsFull() bool {
	return !this.IsEmpty() && (this.rear+1)%this.maxSize == this.front

}

func Main() {
	obj := Constructor(5)
	fmt.Println(obj.EnQueue(1))
	fmt.Println(obj.EnQueue(2))
	fmt.Println(obj.EnQueue(3))
	fmt.Println(obj.EnQueue(4))
	fmt.Println(obj.DeQueue())

	param_3 := obj.Front()
	fmt.Println(param_3)

	param_4 := obj.Rear()
	fmt.Println(param_4)

	param_5 := obj.IsEmpty()
	fmt.Println(param_5)

	param_6 := obj.IsFull()
	fmt.Println(param_6)

}
