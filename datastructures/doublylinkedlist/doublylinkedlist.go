package dll

import (
	"fmt"
)

// Node is a node in the doubly linked list
type Node struct {
	value int
	next  *Node
	prev  *Node
}

// DLinkedList is the doubly linked list
type DLinkedList struct {
	size int
	head *Node
	tail *Node
}

// newNode creates a new node and returns it.
func newNode(value int) Node {
	newNode := Node{value: value}
	return newNode
}

// Insert a new value to the linked list
func (list *DLinkedList) Insert(value int) {
	node := newNode(value)
	if list.size == 0 {
		list.head = &node
		list.tail = &node
	} else {
		list.tail.next = &node
		node.prev = list.tail
		list.tail = &node
	}
	list.size++
}

// Print the list
func (list *DLinkedList) Print() {
	node := list.head
	for {
		if node != nil {
			fmt.Println(node.value)
			if node.next == nil {
				break
			}
			node = node.next
		}
	}
}
