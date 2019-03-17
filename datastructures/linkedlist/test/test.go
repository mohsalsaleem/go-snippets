package main

import (
	"fmt"

	sll "github.com/mohsalsaleem/go-snippets/datastructures/linkedlist"
)

func main() {
	list := sll.SLinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(2)
	list.Insert(4)
	list.Print()

	fmt.Println("Does list have 2? - ", list.Contains(2))

	list.Delete(2)
	list.Delete(2)

	fmt.Println("Does list have 2? - ", list.Contains(2))

	list.Print()
	fmt.Println("Does list have 1? - ", list.Contains(1))
	list.Delete(1)
	fmt.Println("Does list have 1? - ", list.Contains(1))
	list.Delete(4)
	list.Print()
}
