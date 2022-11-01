package main

import (
	"fmt"
	"sort"
)

type SortPeople struct {
	Name string
	Id   int
	Sex  int
}

func printPeople(s []SortPeople) {
	for i, item := range s {
		fmt.Printf("%d:%6s %4d %1d\n", i, item.Name, item.Id, item.Sex)
	}
}

type PeopleHub []SortPeople

func (p PeopleHub) Len() int { return len(p) }
func (p PeopleHub) Less(i, j int) bool {
	if p[i].Name == p[j].Name {
		return p[i].Id < p[j].Id
	}
	return p[i].Name < p[j].Name
}
func (p PeopleHub) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

var peoples = []SortPeople{
	{"apple", 2022, 0},
	{"cat", 10, 0},
	{"apple", 2021, 1},
	{"banana", 10, 0},
}

func main() {
	sort.Sort(PeopleHub(peoples))
	printPeople(peoples)
}
