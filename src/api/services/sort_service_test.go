package services

import (
	"testing"
	"fmt"
)

func TestSort(t *testing.T){
	// Init
	elements := []int{9,7,5,3,1,2,4,6,8,0}
	fmt.Println(elements)
	
	//Execution
	Sort(elements)
	fmt.Println(elements)
		
	//Validation
	if elements[0] != 9 {
		t.Error("first element should be 9")
	}
	
	if elements[len(elements)-1] != 0 {
		t.Error("first element should be 0")
	}
}

