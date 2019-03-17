package main

import (
	sll "github.com/mohsalsaleem/go-snippets/datastructures/linkedlist"
)

func main() {
	list := sll.SLinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(4)
	list.Print()
}
