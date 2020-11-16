package service

import (
	bubblesort "github.com/ravikanthdba/golang-testing/basic-integration-testing/bubblesort"
)

// Sort - is a service that calls bubble sort
func Sort(elements []int) {
	bubblesort.BubbleSort(elements)
}
