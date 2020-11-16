package service

import (
	"testing"
)

func getElements(n int) []int {
	var result = make([]int, n)
	for i := n; i > 0; i-- {
		result = append(result, i)
	}
	return result
}

func BenchmarkSortService(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		SortService(elements)
	}
}
