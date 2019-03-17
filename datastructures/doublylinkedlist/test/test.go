package main

import (
	dll "github.com/mohsalsaleem/go-snippets/datastructures/doublylinkedlist"
)

func main() {
	list := dll.DLinkedList{}
	list.Insert(1)
	list.Insert(3)
	list.Print()
}
