// Blackbox testing, since Sort is only the public function in bubble sort, we can only test the public functions / public variables with black box.

package testing

import (
	"testing"

	services "github.com/ravikanthdba/golang-testing/basic-integration-testing/service"
)

func TestSort(t *testing.T) {
	elements := []int{1, 9, 2, 4, 3, 8, 7, 6, 5}
	services.Sort(elements)
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
