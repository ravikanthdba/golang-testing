package asserttesting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetElements(t *testing.T) {
	elements := GetElements(25)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 25, len(elements))
	assert.EqualValues(t, 25, elements[0])
	assert.EqualValues(t, 1, elements[len(elements)-1])
}

func TestBubbleSort(t *testing.T) {
	elements := GetElements(25)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 25, len(elements))
	assert.EqualValues(t, 25, elements[0])
	assert.EqualValues(t, 1, elements[len(elements)-1])
	fmt.Println(elements)
	BubbleSort(elements)
	fmt.Println(elements)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 25, len(elements))
	assert.EqualValues(t, 1, elements[0])
	assert.EqualValues(t, 25, elements[len(elements)-1])
}

func TestSortingUsingPackage(t *testing.T) {
	elements := GetElements(25)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 25, len(elements))
	assert.EqualValues(t, 25, elements[0])
	assert.EqualValues(t, 1, elements[len(elements)-1])
	fmt.Println(elements)
	SortingUsingPackage(elements)
	fmt.Println(elements)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 25, len(elements))
	assert.EqualValues(t, 1, elements[0])
	assert.EqualValues(t, 25, elements[len(elements)-1])
}
