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

// Contains function checks if the given value is present in the list
func (list *SLinkedList) Contains(value int) bool {
	node := list.head
	for {
		if node != nil {
			if node.value == value {
				return true
			}
			node = node.next
		} else {
			break
		}
	}
	return false
}

// Delete function deletes the first occurance of the node by value
func (list *SLinkedList) Delete(value int) {
	if list.size == 0 {
		return
	}

	if list.head.value == value && list.size == 1 {
		list.head = nil
		list.tail = nil
		list.size = 0
		return
	}

	if list.head.value == value {
		list.head = list.head.next
		list.size--
		return
	}

	curr := list.head.next
	prev := list.head

	for {
		if curr.value == value {
			// Case for tail
			if curr.next == nil {
				prev.next = nil
				list.tail = prev
			} else {
				prev.next = curr.next
			}
			curr = nil
			list.size--
			break
		} else {
			// If the current node's next is not nil, move forward
			// Else, break the loop
			if curr.next != nil {
				curr = curr.next
				prev = prev.next
			} else {
				break
			}
		}
	}
}

// Print function prints the list
func (list *SLinkedList) Print() {
	fmt.Println("==============")
	if list.size == 0 {
		fmt.Println("= Empty List =")
	}
	node := list.head
	for {
		if node != nil {
			fmt.Println(node.value)
			node = node.next
		} else {
			break
		}
	}
}
