package sort

import (
	"testing"
	"fmt"
)

// BubleSort function
func TestBubleSort(t *testing.T) {
	// Init
	elements := []int{9,7,5,3,1,2,4,6,8,0}
	fmt.Println(elements)

	//Execution
	BubleSort(elements)
	fmt.Println(elements)
	
	//Validation
	if elements[0] != 9 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("first element should be 0")
	}
}

// BubleSort function - proof that coverage is not a good metric
func TestBubleSortAlready(t *testing.T) {
	// Init
	elements := []int{9,7,5,3}

	//Execution
	BubleSort(elements)
	
	
}