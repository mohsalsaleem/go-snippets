package sll

import "fmt"

// Node is the single node in the singly linked list
type Node struct {
	value int
	next  *Node
}

// SLinkedList is the singly linked list
type SLinkedList struct {
	head *Node
	tail *Node
	size int
}

// NewNode creates a new node with a value and return it.
func newNode(value int) Node {
	newNode := Node{}
	newNode.value = value
	newNode.next = nil
	return newNode
}

// Insert function inserts a value in the list
func (list *SLinkedList) Insert(value int) {
	node := newNode(value)
	if list.size == 0 {
		list.head = &node
		list.tail = &node
	} else {
		list.tail.next = &node
		list.tail = list.tail.next
	}
	list.size++
}

// TODO: Implement delete and constains operations

// Print function prints the list
func (list *SLinkedList) Print() {
	var node = list.head
	for {
		if node != nil {
			fmt.Println(node.value)
			node = node.next
		} else {
			break
		}
	}
}
