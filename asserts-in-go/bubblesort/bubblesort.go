package bubblesort

import "sort"

//BubbleSort - Takes an array of unsorted elements and sorts them into ascending order
func BubbleSort(elements []int) []int {
	working := true
	for working {
		working = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				working = true
			}
		}
	}
	return elements
}

//SortingUsingPackage - Takes an array of unsorted elements and sorts them by using the sort built-in package
func SortingUsingPackage(elements []int) []int {
	sort.Ints(elements)
	return elements
}

//GetElements - Takes an integer and creates an array of integers from n to 1 in descending order
func GetElements(n int) []int {
	var elements []int
	for i := n; i > 0; i-- {
		elements = append(elements, i)
	}
	return elements
}
