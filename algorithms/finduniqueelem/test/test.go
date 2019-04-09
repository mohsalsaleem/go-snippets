package main

import (
	"fmt"

	finduniqueelem "github.com/mohsalsaleem/go-snippets/algorithms/finduniqueelem"
)

func main() {
	elems := []int{1, 1, 4, 3, 4}
	fmt.Println("Unique Elements: ", finduniqueelem.FindUniqueElement(elems...))
}

//go build github.com/mohsalsaleem/go-snippets/algorithms/finduniqueele
//go install github.com/mohsalsaleem/go-snippets/algorithms/finduniqueelem/test
