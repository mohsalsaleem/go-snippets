package finduniqueelem

// FindUniqueElement finds and prints the unique elements in an array
func FindUniqueElement(elems ...int) []int {
	countMap := make(map[int]int)

	// for _, elem := range elems {
	// 	countMap[elemt] = 0
	// }

	for _, elem := range elems {
		countMap[elem]++
	}

	uniqueElements := make([]int, 0)
	for _, elem := range elems {
		if countMap[elem] == 1 {
			uniqueElements = append(uniqueElements, elem)
		}
	}

	return uniqueElements
}
