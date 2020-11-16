package bubblesort

import (
	"fmt"
	"testing"
	"time"
)

//TestGetElements - Test case for GetElements - takes an integer and prints out all elements from the input integer to 1 in descending order
func TestGetElements(t *testing.T) {
	timeOutChannel := make(chan bool, 1)
	var elements []int
	go func() {
		elements = GetElements(25)
		timeOutChannel <- false
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		timeOutChannel <- true
	}()

	if <-timeOutChannel {
		t.Errorf("Get Elements has failed due to time out error of 500ms, test failed\n")
	} else {
		t.Logf("elements returned without time out error, test passed.\n")
	}

	if len(elements) != 25 {
		t.Errorf("Expected 25 elements, but received only: %d, test failed\n", len(elements))
	} else {
		t.Logf("Expected 25 elements and received: %d, test passed\n", len(elements))
	}

	if elements[0] == 25 {
		t.Logf("Unsorted Order's first element should be 25 and received %d, test passed\n", elements[0])
	} else {
		t.Errorf("Unsorted Order's first element should be 25 and received %d, test failed\n", elements[0])
	}

	if elements[len(elements)-1] == 1 {
		t.Logf("Unsorted Order's last element should be 1 and received %d, test passed\n", elements[len(elements)-1])
	} else {
		t.Errorf("Unsorted Order's last element should be 1 and received %d, test failed\n", elements[len(elements)-1])
	}

}

//TestBubbleSort - Unit tests the BubbleSort function and GetElements function without the assert statement
//1. Checks for the number of elements after the GetElements Function
//2. Checks for the first and last element in unsorted array
//3. Sorts the Array using Bubble Sort
//4. Checks number of elements after the sorting
//5. Checks first and Last elements after the sorting
func TestBubbleSort(t *testing.T) {
	elements := GetElements(10)
	timeOutError := make(chan bool, 1)

	if len(elements) != 10 {
		t.Errorf("Expected 10 elements and got %d elements, test failed\n", len(elements))
	} else {
		t.Logf("Expected 10 elements and received %d elements, test passed\n", len(elements))
	}

	if elements[0] == 10 {
		t.Logf("Unsorted Order's first element should be 10 and received %d, test passed\n", elements[0])
	} else {
		t.Errorf("Unsorted Order's first element should be 10 and received %d, test failed\n", elements[0])
	}

	if elements[len(elements)-1] == 1 {
		t.Logf("Unsorted Order's last element should be 1 and received %d, test passed\n", elements[len(elements)-1])
	} else {
		t.Errorf("Unsorted Order's last element should be 1 and received %d, test failed\n", elements[len(elements)-1])
	}

	fmt.Println(elements)
	go func() {
		BubbleSort(elements)
		timeOutError <- false
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		timeOutError <- false
	}()

	if <-timeOutError {
		t.Errorf("Bubble Sort did not complete within 500ms, test failed\n")
	} else {
		t.Logf("Bubble Sort completed, test passed\n")
	}

	fmt.Println(elements)

	if len(elements) != 10 {
		t.Errorf("Expected 10 elements and got %d elements, test failed\n", len(elements))
	} else {
		t.Logf("Expected 10 elements and received %d elements, test passed\n", len(elements))
	}

	if elements[0] == 1 {
		t.Logf("Sorted Order's first element should be 9 and received %d, test passed\n", elements[0])
	} else {
		t.Errorf("Sorted Order's first element should be 9 and received %d, test failed\n", elements[0])
	}

	if elements[len(elements)-1] == 10 {
		t.Logf("Sorted Order's last element should be 10 and received %d, test passed\n", elements[len(elements)-1])
	} else {
		t.Errorf("Sorted Order's last element should be 10 and received %d, test failed\n", elements[len(elements)-1])
	}
}

//TestSortingUsingPackage - Unit tests the Package sorting function and GetElements function without the assert statement
//1. Checks for the number of elements after the GetElements Function
//2. Checks for the first and last element in unsorted array
//3. Sorts the Array using Package sorting Sort
//4. Checks number of elements after the sorting
//5. Checks first and Last elements after the sorting
func TestSortingUsingPackage(t *testing.T) {
	elements := GetElements(10)
	timeOutError := make(chan bool, 1)
	if len(elements) != 10 {
		t.Errorf("Expected 10 elements and got %d elements, test failed\n", len(elements))
	} else {
		t.Logf("Expected 10 elements and received %d elements, test passed\n", len(elements))
	}

	if elements[0] == 10 {
		t.Logf("Unsorted Order's first element should be 10 and received %d, test passed\n", elements[0])
	} else {
		t.Errorf("Unsorted Order's first element should be 10 and received %d, test failed\n", elements[0])
	}

	if elements[len(elements)-1] == 1 {
		t.Logf("Unsorted Order's last element should be 1 and received %d, test passed\n", elements[len(elements)-1])
	} else {
		t.Errorf("Unsorted Order's last element should be 1 and received %d, test failed\n", elements[len(elements)-1])
	}

	fmt.Println(elements)
	go func() {
		SortingUsingPackage(elements)
		timeOutError <- false
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		timeOutError <- true
	}()

	if <-timeOutError {
		t.Errorf("Sorting by using the package did not complete in 500ms, test failed\n")
	} else {
		t.Logf("Sorting completed, test passed\n")
	}
	fmt.Println(elements)

	if len(elements) != 10 {
		t.Errorf("Expected 10 elements and got %d elements, test failed\n", len(elements))
	} else {
		t.Logf("Expected 10 elements and received %d elements, test passed\n", len(elements))
	}

	if elements[0] == 1 {
		t.Logf("Sorted Order's first element should be 9 and received %d, test passed\n", elements[0])
	} else {
		t.Errorf("Sorted Order's first element should be 9 and received %d, test failed\n", elements[0])
	}

	if elements[len(elements)-1] == 10 {
		t.Logf("Sorted Order's last element should be 10 and received %d, test passed\n", elements[len(elements)-1])
	} else {
		t.Errorf("Sorted Order's last element should be 10 and received %d, test failed\n", elements[len(elements)-1])
	}
}
