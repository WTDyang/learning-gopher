package main

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	var testPeo = []SortPeople{
		{"apple", 2022, 0},
		{"cat", 10, 0},
		{"apple", 2021, 1},
		{"banana", 10, 0},
	}
	sort.Sort(PeopleHub(testPeo))
	if !sort.IsSorted(PeopleHub(testPeo)) {
		t.Error("the slice is not sorted")
	}
}
