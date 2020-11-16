package sort

import (
	"testing"
	"time"
)

func getElements(n int) []int {
	var result = make([]int, n)
	for i := n; i > 0; i-- {
		result = append(result, i)
	}
	return result
}

func TestBubbleSort(t *testing.T) {
	elements := getElements(100)
	timerChannel := make(chan bool, 1)

	go func() {
		BubbleSort(elements)
		timerChannel <- false
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		timerChannel <- true
	}()

	if <-timerChannel {
		t.Fatalf("Test case failed, as bubble sort failed to execute in 500ms")
		return
	}

	if elements[0] == 0 {
		t.Log("Test cases passed")
	} else {
		t.Fatalf("Expected %d, but got %d", 1, elements[0])
	}
}
