package main

import (
	"fmt"

	dll "github.com/mohsalsaleem/go-snippets/datastructures/doublylinkedlist"
)

func main() {
	list := dll.DLinkedList{}

	fmt.Println("Insert 1")
	list.Insert(1)

	fmt.Println("Insert 3")
	list.Insert(3)

	fmt.Println("Insert 4")
	list.Insert(4)

	fmt.Println("Insert 5")
	list.Insert(5)

	fmt.Println("Print list")
	list.Print()

	fmt.Println("Delete 3")
	list.Delete(3)
	fmt.Println("Print list")
	list.Print()

	fmt.Println("Delete 1")
	list.Delete(1)
	fmt.Println("Print list")
	list.Print()

	fmt.Println("Delete 4")
	list.Delete(4)

	fmt.Println("Contains 1? - ", list.Contains(1))

	fmt.Println("Print list")
	list.Print()

	fmt.Println("Insert 1")
	list.Insert(1)

	fmt.Println("Contains 1? - ", list.Contains(1))

	fmt.Println("Print list")
	list.Print()
}
