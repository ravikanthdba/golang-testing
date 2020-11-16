package service

import (
	sortfunction "github.com/ravikanthdba/golang-testing/basic-benchmarking/sorting"
)

//SortService - takes an unsorted array and sorts it
func SortService(elements []int) {
	if len(elements) < 10000 {
		sortfunction.BubbleSort(elements)
		return
	}
	sortfunction.PackageSort(elements)
}
