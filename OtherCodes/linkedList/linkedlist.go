package linkedList

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

func (n *Node) NewNode(i int) {
	n.data = i
	n.next = nil
}

func NewNode(i int) *Node {
	return &Node{i, nil}
}

type SinglyLinkedList struct {
	head *Node
}

func (l *SinglyLinkedList) add(i int) {
	// var temp = new(Node)
	// temp.NewNode(i)
	// var temp *Node
	// temp = NewNode(i)
	temp := &Node{i, nil}

	if l.head == nil {
		l.head = temp
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = temp
	}
}

func (l *SinglyLinkedList) print() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Print("\n")

}

func Main() {
	var l SinglyLinkedList
	l.add(10)
	l.add(9)
	l.add(8)
	l.add(4)
	l.add(1)
	l.print()
	sl := SinglyLinkedList{}
	sl.add(1)
	sl.add(9)
	sl.add(4)
	sl.add(10)
	sl.add(8)
	sl.print()
}
