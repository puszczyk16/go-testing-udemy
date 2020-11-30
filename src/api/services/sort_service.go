package services

import (
	"github.com/puszczyk16/go-testing-udemy/src/api/utils/sort"
)

func Sort(elements []int) {
	if len(elements) <= 10000 {
		sort.BubleSort(elements)
		return
	}
	sort.Sort(elements)
}