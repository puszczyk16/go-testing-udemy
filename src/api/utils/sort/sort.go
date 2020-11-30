package sort

import (
	"sort"
)

// BubleSort function
func BubleSort(elements []int) {
	keepWorking := true
	for keepWorking {
		keepWorking = false
		for i:=0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepWorking = true
			}
		}
	}
}
// Sort function
func Sort(elements []int) {
	sort.Ints(elements)
}