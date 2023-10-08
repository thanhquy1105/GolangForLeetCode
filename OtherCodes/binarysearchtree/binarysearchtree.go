package binarysearchtree

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(root *Node, value int) *Node {
	newNode := &Node{value, nil, nil}
	if root == nil {
		return newNode
	}
	if root.value > value {
		root.left = insert(root.left, value)
	} else {
		root.right = insert(root.right, value)
	}

	// cur := root
	// for cur != nil {
	// 	if value < cur.value {
	// 		if cur.left == nil {
	// 			cur.left = newNode
	// 			break
	// 		} else {
	// 			cur = cur.left
	// 		}
	// 	}
	// 	if value > cur.value {
	// 		if cur.right == nil {
	// 			cur.right = newNode
	// 			break

	// 		} else {
	// 			cur = cur.right
	// 		}
	// 	}
	// }

	return root
}

func search(root *Node, value int) *Node {
	if root == nil {
		return nil
	}
	if root.value == value {
		return root
	}
	if root.value > value {
		return search(root.left, value)
	} else {
		return search(root.right, value)
	}
}

func delete(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	if root.value > value {
		return delete(root.left, value)
	}
	if root.value < value {
		return delete(root.right, value)
	}

	if root.left == nil {
		return root.right
	}
	if root.right == nil {
		return root.left
	}

	newRootParrent := root
	newRoot := root.right

	for newRoot.left != nil {
		newRootParrent = newRoot
		newRoot = newRoot.left
	}

	if newRootParrent == root {
		root.right = newRoot.right
	} else {
		newRootParrent.left = newRoot.right
	}

	root.value = newRoot.value

	return root

}

func show(root *Node) {
	if root == nil {
		return
	}
	show(root.left)
	fmt.Printf("%d ", root.value)
	show(root.right)
}

func show1(root *Node) {
	if root == nil {
		return
	}
	show1(root.right)
	fmt.Printf("%d ", root.value)
	show1(root.left)

}

func showQueue(root *Node) {
	if root == nil {
		return
	}
	queueNode := []Node{}
	queueNode = append(queueNode, Node{root.value, root.left, root.right})

	for len(queueNode) != 0 {
		fmt.Printf("%d ", queueNode[0].value)
		if queueNode[0].left != nil {
			queueNode = append(queueNode, Node{queueNode[0].left.value, queueNode[0].left.left, queueNode[0].left.right})
		}
		if queueNode[0].right != nil {
			queueNode = append(queueNode, Node{queueNode[0].right.value, queueNode[0].right.left, queueNode[0].right.right})
		}
		queueNode = queueNode[1:]
	}
	fmt.Println()
}

func Main() {
	var bst *Node
	/*     50
	    /     \
	   30      70
	  /  \    /  \
	20   40  60   80 */
	bst = insert(bst, 50)
	bst = insert(bst, 30)
	bst = insert(bst, 20)
	bst = insert(bst, 40)
	bst = insert(bst, 70)
	bst = insert(bst, 60)
	bst = insert(bst, 80)
	bst = insert(bst, 55)
	bst = insert(bst, 57)

	show(bst)
	fmt.Println()

	show1(bst)
	fmt.Println()

	showQueue(bst)

	se := search(bst, 60)
	fmt.Println(se)
	se = search(bst, 70)
	fmt.Println(se)
	se = search(bst, 90)
	fmt.Println(se)

	bst = delete(bst, 50)
	show(bst)
	fmt.Println()
	showQueue(bst)

}
