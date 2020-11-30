package services

import (
	"testing"
	/* in thecourse this line blew didn't needed! Witout it service tests don't work! */
	"github.com/puszczyk16/go-testing-udemy/src/api/utils/sort"
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
	if elements[len(elements)-1] != 10000 {
		t.Error("first element should be 10000")
	}
}

func BenchmarkBubbleSort10k(b *testing.B) {
	elements := sort.GetElements(10000)
	for i := 0; i < b.N; i++ {
		sort.BubleSort(elements)
	}
}

func BenchmarkBubbleSort20k(b *testing.B) {
	elements := sort.GetElements(20000)
	for i := 0; i < b.N; i++ {
		sort.Sort(elements)
	}
}
