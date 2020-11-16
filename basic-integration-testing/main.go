package main

import (
	"fmt"

	bubblesort "github.com/ravikanthdba/golang-testing/basic-unit-testing/bubblesort"
)

func main() {
	fmt.Println(bubblesort.BubbleSort([]int{1, 9, 2, 4, 3, 8, 7, 6, 5}))
	fmt.Println(bubblesort.BubbleSort([]int{1200, 1201, 1202, 1204, 103, 18, 765, 600, 15}))
	fmt.Println(bubblesort.BubbleSort([]int{1, 2, 2, 1, 2, 3, 6, 8, 1, 2, 3}))
}
