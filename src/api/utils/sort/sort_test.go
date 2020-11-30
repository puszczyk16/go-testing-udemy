package sort

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// BubleSort function
func TestBubleSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)
	assert.NotNil(t, elements)
	assert.EqualValues(t,10,len(elements))
	assert.EqualValues(t,9,elements[0])
	assert.EqualValues(t,0,elements[9])

	BubleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t,10,len(elements))
	assert.EqualValues(t,9,elements[0])
	assert.EqualValues(t,0,elements[9])

}

func TestSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)
	Sort(elements)
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("first element should be 9")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	//elements := []int{9,7,5,3,1,2,4,6,8,0}
	elements := GetElements(1000)
	for i := 0; i < b.N; i++ {
		BubleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	//elements := []int{9,7,5,3,1,2,4,6,8,0}
	elements := GetElements(1000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}