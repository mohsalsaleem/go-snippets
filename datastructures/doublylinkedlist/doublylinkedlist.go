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

// Delete deletes the first occurrance of a node by alue
func (list *DLinkedList) Delete(value int) {
	node := list.head

	if list.size == 1 {
		if node.value == value {
			node = nil
			list.head = nil
			list.tail = nil
			list.size--
			return
		}
	}

	for {
		if node != nil {
			if node.value == value {
				if node.prev == nil {
					node.next.prev = nil
					list.head = node.next
				} else if node.next == nil {
					node.prev.next = nil
					list.tail = node.prev
				} else {
					node.prev.next = node.next
					node.next.prev = node.prev
					node = nil
				}
				list.size--
				break
			} else {
				node = node.next
			}
		} else {
			break
		}
	}
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

	fmt.Println("Head: ", list.head)
	fmt.Println("Tail: ", list.tail)

	for {
		if node != nil {
			fmt.Println(node.value)
			if node.next == nil {
				break
			}
			node = node.next
		} else {
			break
		}
	}
}
