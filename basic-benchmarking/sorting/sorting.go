package sorting

import "sort"

// BubbleSort - takes list of integers and sorts them in ascending order
func BubbleSort(elements []int) []int {
	for i := 0; i < len(elements)-1; i++ {
		for j := i + 1; j < len(elements); j++ {
			if elements[i] > elements[j] {
				elements[i], elements[j] = elements[j], elements[i]
			}
		}
	}
	return elements
}

//PackageSort - takes an array of unsorted integers and returns the sorted integers, using sort.Intn function
func PackageSort(elements []int) []int {
	sort.Ints(elements)
	return elements
}
