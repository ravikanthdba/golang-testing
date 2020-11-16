package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var array1 = []int{1, 9, 2, 4, 3, 8, 7, 6, 5}
	sortedArray := BubbleSort(array1)
	expectedSorting := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if reflect.DeepEqual(sortedArray, expectedSorting) {
		t.Log("Sorting has been completed, the code is working fine")
	} else {
		t.Fatal("Array is not sorted, the function is not working fine")
	}
}

func TestBubbleSort2(t *testing.T) {
	var array1 = []int{1200, 1201, 1202, 1204, 103, 18, 765, 600, 15}
	sortedArray := BubbleSort(array1)
	expectedSorting := []int{15, 18, 103, 600, 765, 1200, 1201, 1202, 1204}
	if reflect.DeepEqual(sortedArray, expectedSorting) {
		t.Log("Sorting has been completed, the code is working fine")
	} else {
		t.Fatal("Array is not sorted, the function is not working fine")
	}
}

func TestBubbleSort3(t *testing.T) {
	var array1 = []int{1, 2, 2, 1, 2, 3, 6, 8, 1, 2, 3}
	sortedArray := BubbleSort(array1)
	expectedSorting := []int{1, 1, 1, 2, 2, 2, 2, 3, 3, 6, 8}
	if reflect.DeepEqual(sortedArray, expectedSorting) {
		t.Log("Sorting has been completed, the code is working fine")
	} else {
		t.Fatal("Array is not sorted, the function is not working fine")
	}
}
