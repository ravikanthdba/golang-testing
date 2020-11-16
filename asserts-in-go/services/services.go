package services

import (
	"github.com/ravikanthdba/golang-testing/asserts-in-go/asserttesting"
	"github.com/ravikanthdba/golang-testing/asserts-in-go/bubblesort"
)

//SortServiceBubbleSort takes unsorted list of elements, queries the bubblesort package and sorts the elements in ascending order
func SortServiceBubbleSort(elements []int) []int {
	bubblesort.BubbleSort(elements)
	return elements
}

//SortServicePackageSort takes unsorted list of elements, queries the bubblesort package and sorts the elements in ascending order
func SortServicePackageSort(elements []int) []int {
	bubblesort.SortingUsingPackage(elements)
	return elements
}

//SortServiceBubbleSortUsingAssert takes unsorted list of elements, queries the bubblesort package and sorts the elements in ascending order
func SortServiceBubbleSortUsingAssert(elements []int) []int {
	asserttesting.BubbleSort(elements)
	return elements
}

//SortServicePackageSortUsingAssert takes unsorted list of elements, queries the bubblesort package and sorts the elements in ascending order
func SortServicePackageSortUsingAssert(elements []int) []int {
	asserttesting.SortingUsingPackage(elements)
	return elements
}
