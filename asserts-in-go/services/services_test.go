package services

import (
	"testing"

	"github.com/ravikanthdba/golang-testing/asserts-in-go/bubblesort"
	"github.com/stretchr/testify/assert"
)

func TestSortServiceBubbleSort(t *testing.T) {
	elements := bubblesort.GetElements(20)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 20, len(elements))
	assert.EqualValues(t, 20, elements[0])
	assert.EqualValues(t, 1, elements[len(elements)-1])

	SortServiceBubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 20, len(elements))
	assert.EqualValues(t, 1, elements[0])
	assert.EqualValues(t, 20, elements[len(elements)-1])
}

func TestSortServicePackageSort(t *testing.T) {
	elements := bubblesort.GetElements(20)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 20, len(elements))
	assert.EqualValues(t, 20, elements[0])
	assert.EqualValues(t, 1, elements[len(elements)-1])

	SortServicePackageSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 20, len(elements))
	assert.EqualValues(t, 1, elements[0])
	assert.EqualValues(t, 20, elements[len(elements)-1])
}

func TestSortServiceBubbleSortUsingAssert(t *testing.T) {
	elements := bubblesort.GetElements(20)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 20, len(elements))
	assert.EqualValues(t, 20, elements[0])
	assert.EqualValues(t, 1, elements[len(elements)-1])

	SortServiceBubbleSortUsingAssert(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 20, len(elements))
	assert.EqualValues(t, 1, elements[0])
	assert.EqualValues(t, 20, elements[len(elements)-1])
}

func TestSortServicePackageSortUsingAssert(t *testing.T) {
	elements := bubblesort.GetElements(20)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 20, len(elements))
	assert.EqualValues(t, 20, elements[0])
	assert.EqualValues(t, 1, elements[len(elements)-1])

	SortServicePackageSortUsingAssert(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 20, len(elements))
	assert.EqualValues(t, 1, elements[0])
	assert.EqualValues(t, 20, elements[len(elements)-1])
}
