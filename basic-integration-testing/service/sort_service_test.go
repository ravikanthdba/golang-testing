package service

import (
	"testing"
)

func TestSort(t *testing.T) {
	elements := []int{1, 9, 2, 4, 3, 8, 7, 6, 5}
	Sort(elements)
	if elements[0] != 1 {
		t.Fatalf("Expected 1 but received %d", elements[0])
	} else {
		t.Log("Test case has passed")
	}

	if elements[len(elements)-1] == 9 {
		t.Log("Test case has passed")
	} else {
		t.Fatalf("Expected 1 but received %d", elements[len(elements)-1])
	}
}
