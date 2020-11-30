package sort

import (
	"testing"
)

// BubleSort function
func TestBubleSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)
	BubleSort(elements)
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("first element should be 9")
	}
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