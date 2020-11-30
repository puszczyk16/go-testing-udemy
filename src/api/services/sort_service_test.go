package services

import (
	"testing"
)

func TestSort(t *testing.T){
	elements := sort.GetElements(10)
	Sort(elements)
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("first element should be 9")
	}
}

func TestSortMoraThan10000(t *testing.T) {
	elements := sort.GetElements(10001)
	Sort(elements)
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("first element should be 9")
	}
}


/*
func getElements(n int) []int{
	result := make([]int, n)
	j := 0
	for i := n-1; i > 0; i-- {
		result[j] = i
		j++ 
	}
	return result
}
*/