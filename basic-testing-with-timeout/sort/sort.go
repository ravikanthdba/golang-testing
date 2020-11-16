package sort

//BubbleSort - Takes a list of unsorted elements and returns the sorted elements
func BubbleSort(elements []int) []int {
	keepWorking := true
	for keepWorking {
		keepWorking = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepWorking = true
			}
		}
	}
	return elements
}
